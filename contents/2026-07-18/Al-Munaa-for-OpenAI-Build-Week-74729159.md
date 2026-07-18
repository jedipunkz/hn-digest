---
source: "https://devpost.com/software/zeedos-self-hosted-autonomous-ai-operating-system"
hn_url: "https://news.ycombinator.com/item?id=48962403"
title: "Al-Munaa for OpenAI Build Week"
article_title: "AL-MUNAA - Collective Immune System for AI Agents | Devpost"
author: "FARHANALMUTAIRI"
captured_at: "2026-07-18T21:41:06Z"
capture_tool: "hn-digest"
hn_id: 48962403
score: 1
comments: 0
posted_at: "2026-07-18T21:06:38Z"
tags:
  - hacker-news
  - translated
---

# Al-Munaa for OpenAI Build Week

- HN: [48962403](https://news.ycombinator.com/item?id=48962403)
- Source: [devpost.com](https://devpost.com/software/zeedos-self-hosted-autonomous-ai-operating-system)
- Score: 1
- Comments: 0
- Posted: 2026-07-18T21:06:38Z

## Translation

タイトル: OpenAI Build Week の Al-Munaa
記事のタイトル: AL-MUNAA - AI エージェントのための集団免疫システム |開発ポスト
説明: AL-MUNAA - AI エージェント用の集団免疫システム - AI エージェントが署名されたプライバシーに安全な脅威抗体を共有できるようにする開発者ツールです。

記事本文:
サポートされていないブラウザを使用していることが検出されました。お願いします
ブラウザをアップグレードしてください
Internet Explorer 10 以降。
公開ハッカソンに参加する
会社のプライベートハッカソンにアクセスする
開発者エコシステムを成長させ、プラットフォームを宣伝します
組織内のイノベーション、コラボレーション、維持を推進します
ハッカソンの計画と参加に関する洞察
同業者や他の業界リーダーからのインスピレーション
オンラインおよび対面のハッカソンを計画するためのベスト プラクティス
今後のイベントとオンデマンド録画
よくある質問とサポートドキュメント
公開ハッカソンに参加する
会社のプライベートハッカソンにアクセスする
開発者エコシステムを成長させ、プラットフォームを宣伝します
組織内のイノベーション、コラボレーション、維持を推進します
ハッカソンの計画と参加に関する洞察
同業者や他の業界リーダーからのインスピレーション
オンラインおよび対面のハッカソンを計画するためのベスト プラクティス
今後のイベントとオンデマンド録画
よくある質問とサポートドキュメント
AL-MUNAA - AI エージェントのための集団免疫システム
AI エージェントが署名されたプライバシー保護された脅威抗体を共有できるようにする開発者ツール。これにより、あるエージェントによってブロックされた変異プロンプト インジェクションを、ツールが実行される前に別のエージェントによってブロックできます。
AI エージェントは、ファイルを読み取り、ツールを呼び出し、システム全体で動作できます。この力により、新たな障害モードが作成されます。ドキュメントに隠された間接的なプロンプト インジェクションにより、本来は有用なエージェントがシークレットを読み取って攻撃者に送信するように仕向けられる可能性があります。従来のプロンプト フィルターはローカルで脆弱です。 AL-MUNAA は別の質問をします。エージェントが免疫記憶を開発し、元のプライベート プロンプトや秘密を共有せずに他の信頼できるエージェントに警告できたらどうなるでしょうか?
AL-MUNAA は、エージェントの安全のための開発者ツールです。エージェントのワークフローをラップします。

プットおよびメモリ スキャン、アクション ゲート、出力検証、および脅威抗体プロトコル。 1 つのエージェントが攻撃を検出すると、署名されたプライバシーに安全な HMAC フィンガープリントが作成されます。同じ信頼ファミリー内の別のエージェントは、その抗体を検証し、危険なツール呼び出しが実行される前に、変異したバージョンの攻撃をブロックできます。
デモでは、同じ悪意のある Runbook が保護なしで成功します。保護されていないエージェントは読み取りツールとシンク ツールに到達します。 AL-MUNAA を有効にすると、アクション ゲートは危険な呼び出しを実行前にブロックします。次に、2 番目のエージェントは署名された抗体を受け取り、元の秘密を含むプロンプトを見ることなく、改変された攻撃をブロックします。
コアは、決定論的なガードレールと再現可能なテストを備えた Python パッケージです。抗体マッチャーは、Jaccard の類似性とパディング耐性の包含マッチングを組み合わせ、境界のある HMAC スケッチを使用し、異なる信頼ファミリーからインポートされた抗体を拒否します。リポジトリには CLI デモ、キャリブレーション スクリプト、CI マトリックス、インストール可能な GitHub リリースが含まれているため、審査員は最初から再構築することなくテストできます。
Codex と GPT-5.6 の使用方法
GPT-5.6-Sol 上の Codex は、コア プロトコルでテストファースト作業を実行しました。しきい値に近い脆弱性を再現し、封じ込めマッチングを追加し、制限された HMAC スケッチを 256 エントリから 512 エントリに拡張し、トラストファミリーの拒否を強制し、キャリブレーション マトリックスを作成し、インストールされたパッケージのエントリポイントを修復し、証拠を文書化しました。
GPT-5.6 は、OpenAI Responses API を介して製品内で、グレーケースの構造化インテント アナリストとして、また、限定的防御ワクチン ジェネレーターとしても使用されます。決定論的な政策は依然として最終的な権限である。 GPT-5.6 は不確実なケースについてアドバイスし、ハッシュとしてのみ保存される防御的なバリアントを生成します。
コーデックス /フィードバック セッション ID: 019f72

df-5fa7-7cf1-95f4-265467d02099
完全なローカル スイート: 74 のテストに合格しました。
レガシー マッチャー: 3/4 の攻撃が検出され、0/5 の誤検知。
調整されたマッチャー: 4/4 の攻撃が検出され、0/5 の誤検知。
デモ証拠金は 0.0264 から 0.1474 に改善しました。
64 の決定論的ファミリー キーにより、576 のフィクスチャ評価が生成され、フィクスチャ エラーはゼロでした。
ライブ GPT-5.6 Sol アクション ゲート ベンチマーク: 保護されていない読み取り = 1/シンク = 1;保護されたガードブロック=1/読み取り=0/シンク=0。
ライブ製品内 GPT-5.6 パス: 信頼度 0.99 のグレーケース判定ブロック、およびハッシュとしてのみ保存された 2 つの防御的ワクチンバリアント。
難しいのは、証拠を正直に保つことでした。キャリブレーション マトリックスは意図的に小さく合成されているため、AL-MUNAA は本番環境全体のリコール、誤検知ゼロ、またはあらゆるセマンティック書き換えからの保護を主張していません。現在のバージョンは、焦点を絞った垂直スライスです。アイデアを実証するのに十分な強度があり、審査員がテストできるほど十分に狭いです。
次のステップは、より大規模な現実世界の攻撃コーパス、一般的なエージェント ランタイム用のフレームワーク アダプター、チームを対象としたトラストファミリー管理、およびエージェント フリート全体にわたる抗体伝播を表示するためのダッシュボードです。
リポジトリ: https://github.com/Farhanward/al-munaa
リリース: https://github.com/Farhanward/al-munaa/releases/tag/v0.1.1
ポートフォリオ: https://portfolio.carbonflows.store/?v=farhan-almutairi&lang=ar
ファルハン・アルムタイリ
創設者 @ CarbonFlow — ローカル LLM および RAG を使用して、プライバシー最優先のセルフホスト型 AI システム、自律エージェント、AI セキュリティ ツールを構築しています。
ファルハン・アルムタイリ
更新情報を投稿しました
—
2026 年 7 月 18 日 01:56 PM EDT
AL-MUNAA は、開発者ツール トラックの OpenAI Build Week に提出されました。
このプロジェクトは、AI エージェントの集団免疫システムを実証します。1 つのエージェントがプロンプト インジェクション攻撃を検出し、署名されたプライバシーに安全な脅威抗体を作成し、別の信頼できるエージェントがブロックします。

危険なツールが実行される前に、変異したバージョンをロックします。
デモビデオ: https://youtu.be/mlAxp2UJaFg
GitHub: https://github.com/Farhanward/al-munaa
LinkedIn の発表: AL-MUNAA は、開発者ツール トラックの OpenAI Build Week に提出されました。
このプロジェクトは、AI エージェントの集団免疫システムを実証します。1 つのエージェントがプロンプト インジェクション攻撃を検出し、署名されたプライバシーに安全な脅威抗体を作成し、別の信頼できるエージェントが危険なツールが実行される前に変異バージョンをブロックできます。
デモビデオ: https://youtu.be/mlAxp2UJaFg
GitHub: https://github.com/Farhanward/al-munaa
LinkedIn のお知らせ: https://www.linkedin.com/feed/update/urn:li:share:7484304627665055744
Devpost でこのプロジェクトをレビューして、コメントを残して、共感していただければ、心から感謝します。
OpenAI #OpenAIBuildWeek #Codex #GPT56 #AIAgents #AISafety #CyberSecurity #PromptInjection #DeveloperTools #BuildInPublichttps://www.linkedin.com/feed/update/urn:li:share:7484304627665055744
Devpost でこのプロジェクトをレビューして、コメントを残して、共感していただければ、心から感謝します。
OpenAI #OpenAIBuildWeek #Codex #GPT56 #AIAgents #AISafety #CyberSecurity #PromptInjection #DeveloperTools #BuildInPublic
ログイン
または
Devpost にサインアップする
会話に参加します。
ファルハン・アルムタイリ
このプロジェクトを始めました
—
2026 年 7 月 18 日 01:28 PM EDT
コメントにフィードバックを残してください!
ログイン
または
Devpost にサインアップする
会話に参加します。

## Original Extract

AL-MUNAA - Collective Immune System for AI Agents - A developer tool that lets AI agents share signed, privacy-safe threat antibodies so a mutated ...

We've detected that you are using an unsupported browser. Please
upgrade your browser
to Internet Explorer 10 or higher.
Participate in our public hackathons
Access your company's private hackathons
Grow your developer ecosystem and promote your platform
Drive innovation, collaboration, and retention within your organization
Insights into hackathon planning and participation
Inspiration from peers and other industry leaders
Best practices for planning online and in-person hackathons
Upcoming events and on-demand recordings
Common questions and support documentation
Participate in our public hackathons
Access your company's private hackathons
Grow your developer ecosystem and promote your platform
Drive innovation, collaboration, and retention within your organization
Insights into hackathon planning and participation
Inspiration from peers and other industry leaders
Best practices for planning online and in-person hackathons
Upcoming events and on-demand recordings
Common questions and support documentation
AL-MUNAA - Collective Immune System for AI Agents
A developer tool that lets AI agents share signed, privacy-safe threat antibodies so a mutated prompt-injection blocked by one agent can be blocked by another before tools execute.
AI agents can read files, call tools, and act across systems. That power creates a new failure mode: an indirect prompt injection hidden in a document can convince an otherwise useful agent to read secrets and send them to an attacker. Traditional prompt filters are local and brittle. AL-MUNAA asks a different question: what if agents could develop an immune memory and warn other trusted agents without sharing the original private prompt or secret?
AL-MUNAA is a developer tool for agent safety. It wraps agent workflows with input and memory scanning, an action gate, output verification, and a Threat Antibody Protocol. When one agent detects an attack, it creates a signed, privacy-safe HMAC fingerprint. Another agent in the same trust family can verify that antibody and block a mutated version of the attack before any dangerous tool call executes.
In the demo, the same malicious runbook succeeds without protection: the unprotected agent reaches a read tool and a sink tool. With AL-MUNAA enabled, the action gate blocks the dangerous call before execution. A second agent then receives a signed antibody and blocks a modified attack without seeing the original secret-bearing prompt.
The core is a Python package with deterministic guardrails and reproducible tests. The antibody matcher combines Jaccard similarity with padding-resistant containment matching, uses bounded HMAC sketches, and rejects antibodies imported from a different trust family. The repository includes a CLI demo, calibration script, CI matrix, and an installable GitHub release so judges can test it without rebuilding from scratch.
How Codex and GPT-5.6 were used
Codex on GPT-5.6-Sol performed test-first work on the core protocol. It reproduced a near-threshold weakness, added containment matching, expanded the bounded HMAC sketch from 256 to 512 entries, enforced trust-family rejection, created a calibration matrix, repaired the installed-package entrypoint, and documented the evidence.
GPT-5.6 is also used inside the product through the OpenAI Responses API as a structured intent analyst for gray cases and as a bounded defensive vaccine generator. Deterministic policy remains the final authority; GPT-5.6 advises on uncertain cases and produces defensive variants that are stored only as hashes.
Codex /feedback Session ID: 019f72df-5fa7-7cf1-95f4-265467d02099
Full local suite: 74 tests passed.
Legacy matcher: 3/4 attacks detected, 0/5 false positives.
Calibrated matcher: 4/4 attacks detected, 0/5 false positives.
Demo margin improved from 0.0264 to 0.1474.
64 deterministic family keys produced 576 fixture evaluations with zero fixture errors.
Live GPT-5.6 Sol action-gate benchmark: unprotected read=1/sink=1; protected guard_blocks=1/read=0/sink=0.
Live in-product GPT-5.6 paths: gray-case verdict block at confidence 0.99, plus two defensive vaccine variants stored only as hashes.
The hard part was keeping the evidence honest. The calibration matrix is intentionally small and synthetic, so AL-MUNAA does not claim production-wide recall, zero false positives, or protection from every semantic rewrite. The current version is a focused vertical slice: strong enough to demonstrate the idea, narrow enough to be testable by judges.
Next steps are larger real-world attack corpora, framework adapters for popular agent runtimes, team-scoped trust-family management, and a dashboard for viewing antibody propagation across agent fleets.
Repository: https://github.com/Farhanward/al-munaa
Release: https://github.com/Farhanward/al-munaa/releases/tag/v0.1.1
Portfolio: https://portfolio.carbonflows.store/?v=farhan-almutairi&lang=ar
Farhan Almutairi
Founder @ CarbonFlow — building privacy-first, self-hosted AI systems, autonomous agents, and AI security tooling with local LLMs & RAG.
Farhan Almutairi
posted an update
—
Jul 18, 2026 01:56 PM EDT
AL-MUNAA has been submitted to OpenAI Build Week in the Developer Tools track.
The project demonstrates a collective immune system for AI agents: one agent detects a prompt-injection attack, creates a signed privacy-safe threat antibody, and another trusted agent can block a mutated version before dangerous tools execute.
Demo video: https://youtu.be/mlAxp2UJaFg
GitHub: https://github.com/Farhanward/al-munaa
LinkedIn announcement: AL-MUNAA has been submitted to OpenAI Build Week in the Developer Tools track.
The project demonstrates a collective immune system for AI agents: one agent detects a prompt-injection attack, creates a signed privacy-safe threat antibody, and another trusted agent can block a mutated version before dangerous tools execute.
Demo video: https://youtu.be/mlAxp2UJaFg
GitHub: https://github.com/Farhanward/al-munaa
LinkedIn announcement: https://www.linkedin.com/feed/update/urn:li:share:7484304627665055744
I would be genuinely grateful if you review the project here on Devpost, leave a comment, and like it if it resonates with you.
OpenAI #OpenAIBuildWeek #Codex #GPT56 #AIAgents #AISafety #CyberSecurity #PromptInjection #DeveloperTools #BuildInPublichttps:// www.linkedin.com/feed/update/urn:li:share:7484304627665055744
I would be genuinely grateful if you review the project here on Devpost, leave a comment, and like it if it resonates with you.
OpenAI #OpenAIBuildWeek #Codex #GPT56 #AIAgents #AISafety #CyberSecurity #PromptInjection #DeveloperTools #BuildInPublic
Log in
or
sign up for Devpost
to join the conversation.
Farhan Almutairi
started this project
—
Jul 18, 2026 01:28 PM EDT
Leave feedback in the comments!
Log in
or
sign up for Devpost
to join the conversation.
