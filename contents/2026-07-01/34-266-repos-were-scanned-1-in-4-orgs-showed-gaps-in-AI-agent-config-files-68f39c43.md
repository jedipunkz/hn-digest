---
source: "https://blog.codacy.com/we-scanned-34266-repos.-1-in-4-orgs-showed-gaps-in-ai-agent-config-files"
hn_url: "https://news.ycombinator.com/item?id=48744080"
title: "34,266 repos were scanned: 1 in 4 orgs showed gaps in AI agent config files"
article_title: "We Scanned 34,266 Repos. 1 in 4 Orgs Showed Gaps In AI Agent Config Files"
author: "claudiacsf"
captured_at: "2026-07-01T09:05:43Z"
capture_tool: "hn-digest"
hn_id: 48744080
score: 1
comments: 0
posted_at: "2026-07-01T09:02:40Z"
tags:
  - hacker-news
  - translated
---

# 34,266 repos were scanned: 1 in 4 orgs showed gaps in AI agent config files

- HN: [48744080](https://news.ycombinator.com/item?id=48744080)
- Source: [blog.codacy.com](https://blog.codacy.com/we-scanned-34266-repos.-1-in-4-orgs-showed-gaps-in-ai-agent-config-files)
- Score: 1
- Comments: 0
- Posted: 2026-07-01T09:02:40Z

## Translation

タイトル: 34,266 リポジトリがスキャンされました: 4 組織のうち 1 組織で AI エージェント構成ファイルにギャップが見られました
記事のタイトル: 34,266 リポジトリをスキャンしました。 4 組織に 1 組織で AI エージェント設定ファイルにギャップがあることが判明
説明: AI エージェント構成ファイルのセキュリティと明確さに対する重要なニーズを発見し、より安全なコーディングのためのベスト プラクティスを適用する方法を学びます。

記事本文:
34,266 件のリポジトリをスキャンしました。 4 組織に 1 組織で AI エージェント設定ファイルにギャップがあることが判明
チームは、Claude Code、Cursor、GitHub Copilot などの AI コーディング アシスタントを使用してコードを出荷しており、リポジトリ レベルの指示ファイルによってガイドされることが増えています。問題は、ほとんどのチームがこれらのファイルを運用レベルの構成ではなく、非公式のドキュメントのように扱っていることです。
私たちは、組織が実際にエージェント指示ファイルをどのように管理しているかを確認するために、34,266 個のリポジトリにわたって AgentLinter を実行しました。調査結果では、あいまいな指示、失敗動作の欠落、アプリケーション コードのレビューに決して合格しないセキュリティ リスクなど、一貫したギャップが明らかになりました。この記事では、私たちが発見したことと、コードベースの残りの部分にすでに適用されている同じ規律をエージェント構成に適用するために必要なことを詳しく説明します。
エージェント設定ファイルにセキュリティ制御が必要な理由
34,266 のリポジトリで見つかったこと
調査結果 1: あいまいさが最も一般的な問題です
調査結果 2: 逃走ハッチの欠如が暴走行為を引き起こす
調査結果 3: セキュリティ問題は量は少ないが重大度は高い
定義されたルールが強制されたルールと同じではない理由
エージェントの指示に対する適切な執行とはどのようなものか
今週エージェント構成ファイルを強化する方法
エージェント設定ファイルにセキュリティ制御が必要な理由
ソフトウェア リポジトリ内の AI エージェント構成ファイルを保護するには、いくつかの基本的なプラクティスが必要になります。つまり、シークレットをハードコーディングする代わりに外部化し、厳格なアクセス制御を適用し、構成の明確さと安全性を検証し、漏洩する可能性のある資格情報を継続的にスキャンします。問題は、ほとんどのチームが、アプリケーション コードに適用するのと同じ規律でエージェント構成ファイルの扱いをまだ開始していないことです。マッキンゼーの調査によると、AI ガバナンスが成熟している組織はわずか 30% 程度であることが判明

レベル3以上。
Claude Code 、Cursor、GitHub Copilot などの AI コーディング エージェントは、リポジトリ レベルの命令ファイルに依存して動作をガイドします。 CLAUDE.md 、 .cursorrules 、 copilot-instructions.md 、または AGENTS.md などの名前に見覚えがあるかもしれません。これらのファイルは、エージェントに、従うべきコーディング標準、避けるべきディレクトリ、エラーの処理方法、および実行できるコマンドを伝えます。
ここでリスクが発生します。エージェント指示ファイルは、エージェントがコードを読み取り、編集し、シェル コマンドを実行し、機密データを操作する方法に影響を与えます。しかし、チームは多くの場合、通常のプル リクエストほど厳密ではなく作成およびレビューを行います。指示ファイルがあいまいであったり、矛盾していたり​​、ハードコーディングされた秘密が含まれている場合、エージェントは実行するすべてのタスクにわたってその弱点を増幅させます。
従来の文書は、意図を推測して無関係な詳細を無視できる人間によって解釈されます。一方、エージェント指示ファイルは、テキストに直接作用する LLM ベースのシステムによって使用されます。人間のドキュメントが判断の指針となります。エージェント構成は実行を形成します。
34,266 のリポジトリで見つかったこと
AI エージェント構成ファイル用のオープンソース静的分析ツールである AgentLinter を、この機能を有効にして 34,266 個のリポジトリにわたって実行しました。スキャンは 1,353 の組織を対象としました。
その結果、チームによるアプリケーション コードの扱い方とエージェントの指示の扱い方の間に明らかなギャップがあることがわかりました。
1,604 のリポジトリで、エージェント構成ファイルに問題のフラグが立てられていました
354 の組織が調査結果を含むリポジトリを少なくとも 1 つ持っていた
わかりやすさと明確さに関する 13,000 件以上の問題
避難ハッチの欠落に関連する約 5,000 件の問題
セキュリティリスクに関連する約1,150の問題
調査結果は、あいまいさ、障害動作の欠落、セキュリティの脆弱性の 3 つのカテゴリに分類されます。各カテゴリー担当者

それぞれ異なる種類のリスクに憤慨していますが、根本的な原因は共通しています。チームは、エージェントにガイドとなるルールに対する本番レベルの強制を行わずに、本番レベルの影響力をエージェントに与えています。
調査結果 1: あいまいさが最も一般的な問題です
私たちが発見した最も頻繁な問題は、不明確な指示でした。スキャンされたリポジトリ全体で、未定義の用語が 7,700 件以上出現しました。その他の一般的な問題には、曖昧な指示や、エージェントが確実に解釈できない複雑すぎる文などが含まれていました。
LLM ベースのエージェントは、不明瞭な言語に対して非常に敏感です。指示に「コードを適切にフォーマットする」または「ベスト プラクティスに従ってください」と書かれている場合、エージェントはそれが何を意味するかを推測する必要があります。異なるエージェント、または異なる実行の同じエージェントであっても、解釈が異なる場合があります。
弱い命令と具体的な命令の違いを考えてみましょう。
弱いバージョンではエージェントに推測を強制しますが、具体的なバージョンではテスト可能なものをエージェントに与えます。
指示があいまいな場合、幻覚的な動作が増え、リポジトリ間で一貫性のないコード変更が増え、避けられる間違いを修正するためのレビュー担当者の労力が増加します。
調査結果 2: 逃走ハッチの欠如が暴走行為を引き起こす
スキャンでは、避難ハッチの紛失が 5,000 件近く発生しました。さらに 2,000 件を超える事例には、重複または矛盾した命令が含まれていました。
脱出ハッチは、何か問題が発生したときにエージェントに何をすべきかを指示する明示的な指示です。従来のソフトウェアには、例外、再試行、タイムアウト、フォールバック、エスカレーション パスなどのエラー処理があります。エージェントのワークフローも同じ概念から恩恵を受けます。明示的な失敗動作がないと、エージェントは失敗した同じコマンドをループしたり、無関係なファイルを編集し続けたり、欠落したコンテキストをでっち上げたり、ツールの失敗後に続行したりする可能性があります。
私たちが観察した一般的なアンチパターンは、チームが悪いところを修正しようとすることです。

より強力な指示を繰り返すことによるエージェントの動作:
指示を繰り返すとコンテキストのサイズが増大し、必ずしも強制力を向上させることなく、動作を推論することが困難になる可能性があります。ルールが重複していて矛盾していると、行動の制御が難しくなり、予測が難しくなります。
より良いパターンでは、エージェントに明示的なフォールバック パスが提供されます。
停止条件: 「テストが同じエラーで 2 回失敗した場合、停止して失敗を要約します。」
欠落コンテキストの処理: 「必要なシークレットまたは資格情報が欠落している場合は、作成しないでください。ユーザーに必要な設定を依頼してください。」
構造の検証: 「リポジトリの構造が指示と一致しない場合は、停止して不一致を報告します。」
破壊的なコマンドの確認: 「コマンドに破壊的なフラグが必要な場合は、実行する前に確認を求めます。」
エスケープ ハッチのないエージェントは、エラー処理のないスクリプトのように動作します。外部の何かがそれを止めるまで、それは続きます。
調査結果 3: セキュリティ問題は量は少ないが重大度は高い
セキュリティ関連の調査結果は量は少ないものの、最もリスクの高いカテゴリを表しています。
449 件の危険なコマンド権限
437 件の PII 暴露リスク
71 件のデータ漏洩の脆弱性
エージェント ファイル内の 21 個のハードコードされたシークレット
エージェント構成ファイルでは、エージェントが読み取り、書き込み、実行、または送信を許可される内容を定義できます。ファイルはマークダウンまたは軽量構成であることが多いため、チームはそれらのファイルをセキュリティ上重要なものとして扱わない場合があります。しかし、それらは次のとおりです。OWASP Top 10 for Agentic Applications は、これらのシステムにおける重大なリスクとして、過剰なエージェンシーとツールの誤用を特定しています。
私たちが見つけた危険なパターンの例:
API キーまたはトークンは命令ファイルに直接コミットされます
マスキングせずにサンプルに貼り付けられたユーザー データ
シェルコマンドを実行するための広すぎる権限
cu などの安全でないコマンド パターン

ＲＬ |バッシュ
独自のコードまたはデータを外部サービスにコピーできるようにする手順
エージェントは危険な指示に繰り返し応じて行動する可能性があります。安全でないルールは、多くのリポジトリにコピーされる可能性があります。エージェント ファイルはドキュメントのように見えるため、レビュー担当者は無視する場合があります。既存のスキャナーは、エージェント構成ファイルを高リスクのポリシー サーフェスとして分類しない可能性があります。
ヒント: エージェント構成ファイルをアプリケーションのセキュリティ境界の一部として扱います。これらは、CI/CD 構成、コードとしてのインフラストラクチャ、環境変数定義と同じカテゴリに属します。
定義されたルールが強制されたルールと同じではない理由
多くのチームはすでにエージェントの動作に対して非公式な期待を抱いています。問題は、期待が README ファイル、内部ドキュメント、チャット メッセージ、コピーされたプロンプト スニペット、およびリポジトリ固有の構成ファイルに分散されていることです。
その結果、おなじみの施行ギャップが生じます。
開発者はそれを異なる解釈をします
レビュー担当者が一貫して問題を捉えていない
CI/CD は標準を確実に適用しない
ガバナンスの成熟度は AI 導入に遅れをとっています。 Gartner は、ガバナンスのギャップは運用上のインシデントが発生した後に初めて明らかになるため、2027 年までに 40% の企業が自律型 AI エージェントを降格または廃止すると予測しています。エージェントの安全性は、完璧なプロンプトを一度作成することよりも、変更が発生するたびに許容可能な指示パターンを強制することに依存します。
エージェントの指示に対する適切な執行とはどのようなものか
エージェント構成ファイルの効果的な適用は、コードの品質とセキュリティの適用と同じ原則に従います。開発者が作業している場所をスキャンし、明確さを品質問題として扱い、明示的な失敗動作を要求し、エージェント構成にセキュリティ ポリシーを適用します。
開発者が作業するエージェント ファイルをスキャンします。 IDE フィードバックにより問題を早期に発見します。 Git フックは高速なローカル フィードバックを提供します。プルリクエスト

チェックを入れるとレビューが可視化されます。 CI/CD の適用により、組織全体の一貫性が確保されます。
曖昧さは品質の問題として扱います。未定義の用語、曖昧な指示、および過度に複雑な文にフラグを立てます。エージェントが一貫して従うことができる具体的な表現を奨励します。
明示的な失敗動作が必要です。避難ハッチが欠けていないか確認してください。失敗が繰り返される場合は停止条件が必要です。不確実なコマンドや安全でないコマンドにはエスカレーション パスが必要です。
矛盾を検出し、肥大化を促します。繰り返される指示や矛盾するルールにフラグを立てます。予測不可能性を高める不必要なコンテキストを削減します。
セキュリティ ポリシーをエージェント構成に適用します。エージェント ファイルに対してシークレット検出を実行します。危険なコマンド権限にフラグを立てます。 PII 露出パターンを検出します。漏洩の可能性のある指示がないか確認してください。コマンドの実行、ネットワーク アクセス、または機密パスへのアクセスを許可する権限を確認する必要があります。
所有権を明確にします。共有エージェント ルールのメンテナを割り当てます。エージェント構成の変更には、CODEOWNERS または同等のレビュー ルーティングを使用します。可能な限り、組織レベルのポリシーをリポジトリ固有のガイダンスから分離してください。
今週エージェント構成ファイルを強化する方法
エージェント設定ファイルのセキュリティと信頼性を向上させたい場合は、次の実用的なチェックリストを参照してください。
リポジトリ全体のすべてのエージェント指示ファイル ( CLAUDE.md 、 .cursorrules 、 AGENTS.md 、 copilot-instructions.md など) をインベントリします。
エージェント構成ファイルをコードレビュー範囲に明示的に追加する
エージェント構成ファイルに対して lint を実行して、あいまいさ、エスケープ ハッチの欠落、およびセキュリティ問題を検出します。
ハードコードされたシークレット、トークン、顧客データ、実稼働サンプルを削除します。
曖昧な指示を具体的なルールに置き換える
失敗したテスト、欠落したコンテキスト、コマンドの失敗、および不確実性に対してエスケープ ハッチを追加します。
重複を削除

または矛盾した指示
危険なシェルコマンドを制限し、破壊的な操作には確認を要求します
PR チェックを追加して、エージェント設定の変更がレビューを回避できないようにする
リポジトリ全体で繰り返し発生する検出結果を追跡して、共有ポリシーの問題を特定します
進むべき道はまっすぐです。より明確な指示を書きます。失敗時の動作を定義します。危険な権限と機密データを削除します。リポジトリ全体で一貫してチェックを実施します。
エージェントの指示が実稼働コードに影響を与える場合、それらは実稼働コードと同じエンジニアリング規律に値します。
接続されているリポジトリ全体で AgentLinter を実行して、AI エージェント構成ファイルのあいまいさ、エスケープ ハッチの欠落、セキュリティ リスクを特定します。
リポジトリを無料でスキャン →
エンジニアリング チームにおける AI ツール導入の背後にある可視性の問題
経営陣はエンジニアリングチームにAI導入を加速するよう求めている。しかし、ほとんどのエンジニアリング リーダーは、どの AI コーディング...という基本的な質問に答えることができません。
2026 年に監査人は AI 生成コードについて何を質問するでしょうか
開発者はすでに AI コーディング ツールを使用しています。監査人が問い始めているのは、チームに AI ポリシーがあるかどうかではなく、それができるかどうかです。
ブログを購読する
最新情報を入手してください

[切り捨てられた]

## Original Extract

Discover the crucial need for security and clarity in AI agent configuration files, and learn how to enforce best practices for safer coding.

We Scanned 34,266 Repos. 1 in 4 Orgs Showed Gaps In AI Agent Config Files
Teams are shipping code with AI coding assistants such as Claude Code, Cursor, and GitHub Copilot, increasingly guided by repository-level instruction files. The problem is that most teams treat these files like informal documentation rather than the production-level configuration they have become.
We ran AgentLinter across 34,266 repositories to see how organizations are actually managing their agent instruction files. The findings reveal a consistent gap: ambiguous instructions, missing failure behavior, and security risks that would never pass review in application code . This article breaks down what we found and what it takes to enforce the same discipline on agent configs that you already apply to the rest of your codebase.
Why agent configuration files require security controls
What we found across 34,266 repositories
Finding 1: Ambiguity is the most common issue
Finding 2: Missing escape hatches create runaway behavior
Finding 3: Security issues are lower volume but higher severity
Why defined rules are not the same as enforced rules
What good enforcement looks like for agent instructions
How to harden agent config files this week
Why agent configuration files require security controls
Securing AI agent configuration files in software repositories comes down to a few core practices: externalize secrets instead of hardcoding them, enforce strict access controls, validate configurations for clarity and safety, and continuously scan for potentially leaked credentials. The problem is that most teams have not yet started treating agent config files with the same discipline they apply to application code. A McKinsey survey found only ~30% of organizations at AI governance maturity level 3 or higher.
AI coding agents like Claude Code , Cursor, and GitHub Copilot rely on repository-level instruction files to guide their behavior. You might recognize names like CLAUDE.md , .cursorrules , copilot-instructions.md , or AGENTS.md . The files tell the agent what coding standards to follow, which directories to avoid, how to handle errors, and what commands it can run.
Here is where the risk comes in. Agent instruction files influence how agents read code, make edits, run shell commands, and interact with sensitive data. Yet teams often author and review them with less rigor than a typical pull request. When an instruction file is vague, contradictory, or contains hardcoded secrets, the agent amplifies that weakness across every task it performs.
Traditional documentation is interpreted by humans who can infer intent and ignore irrelevant details. Agent instruction files, on the other hand, are consumed by LLM-based systems that act on the text directly. Human docs guide judgment. Agent configs shape execution.
What we found across 34,266 repositories
We ran AgentLinter, an open-source static analysis tool for AI agent configuration files , across 34,266 repositories with the feature enabled. The scan covered 1,353 organizations.
The results showed a clear gap between how teams treat application code and how they treat agent instructions:
1,604 repositories had issues flagged in their agent config files
354 organizations had at least one repository with findings
Over 13,000 issues related to comprehensibility and clarity
Nearly 5,000 issues related to missing escape hatches
Approximately 1,150 issues related to security risks
The findings break down into three categories: ambiguity, missing failure behavior, and security vulnerabilities. Each category represents a different kind of risk, but they share a common root cause. Teams are giving agents production-level influence without production-level enforcement over the rules that guide them.
Finding 1: Ambiguity is the most common issue
The most frequent problem we found was unclear instructions. More than 7,700 instances of undefined terms appeared across the scanned repositories. Other common issues included vague directives and overly complex sentences that an agent cannot reliably interpret.
LLM-based agents are highly sensitive to unclear language. When an instruction says "format code properly" or "follow best practices," the agent has to infer what that means. Different agents, or even the same agent on different runs, may interpret it differently.
Consider the difference between weak and concrete instructions:
The weak versions force the agent to guess, while the concrete versions give it something testable to follow.
When instructions are ambiguous, you get more hallucinated behavior, more inconsistent code changes across repositories, and more reviewer effort to correct avoidable mistakes.
Finding 2: Missing escape hatches create runaway behavior
Nearly 5,000 occurrences of missing escape hatches appeared in the scan. Another 2,000+ instances involved duplicate or contradictory instructions.
An escape hatch is an explicit instruction that tells the agent what to do when something goes wrong. Traditional software has error handling: exceptions, retries, timeouts, fallbacks, escalation paths. Agent workflows benefit from the same concept. Without explicit failure behavior, agents may loop on the same failing command, keep editing unrelated files, invent missing context, or continue after a tool failure.
A common anti-pattern we observed is teams trying to fix bad agent behavior by repeating stronger instructions:
Repeating instructions increases context size and can make behavior harder to reason about without necessarily improving enforceability. Duplicate and contradictory rules make behavior less predictable, not more controlled.
Better patterns give agents explicit fallback paths:
Stop conditions: "If tests fail twice with the same error, stop and summarize the failure."
Missing context handling: "If a required secret or credential is missing, do not invent one. Ask the user for the required setup."
Structure validation: "If the repository structure does not match the instructions, stop and report the mismatch."
Destructive command confirmation: "If a command requires destructive flags, ask for confirmation before running it."
An agent without an escape hatch behaves like a script without error handling. It will keep going until something external stops it.
Finding 3: Security issues are lower volume but higher severity
Security-related findings were smaller in volume but represent the highest-risk category:
449 instances of dangerous command permissions
437 instances of PII exposure risks
71 data exfiltration vulnerabilities
21 hardcoded secrets in agent files
Agent configuration files can define what the agent is allowed to read, write, run, or transmit. Because the files are often Markdown or lightweight configuration, teams may not treat them as security-sensitive. But they are: the OWASP Top 10 for Agentic Applications identifies excessive agency and tool misuse as critical risks in these systems.
Examples of risky patterns we found:
API keys or tokens committed directly into instruction files
User data pasted into examples without masking
Overly broad permission to run shell commands
Unsafe command patterns such as curl | bash
Instructions that allow copying proprietary code or data into external services
The agent may act on risky instructions repeatedly. Unsafe rules can be copied across many repositories. Reviewers may ignore agent files because they look like documentation. Existing scanners may not classify agent config files as high-risk policy surfaces.
Tip: Treat agent config files as part of the application security boundary. They belong in the same category as CI/CD configuration, infrastructure-as-code, and environment variable definitions.
Why defined rules are not the same as enforced rules
Many teams already have informal expectations for agent behavior. The issue is that expectations are scattered across README files, internal docs, chat messages, copied prompt snippets, and repository-specific config files.
The result is a familiar enforcement gap:
Developers interpret it differently
Reviewers catch issues inconsistently
CI/CD does not reliably enforce the standard
Governance maturity is lagging behind AI adoption. Gartner predicts that, by 2027, 40% of enterprises will demote or decommission autonomous AI agents because governance gaps will only become apparent after production incidents. Agent safety depends less on writing a perfect prompt once and more on enforcing acceptable instruction patterns wherever changes happen.
What good enforcement looks like for agent instructions
Effective enforcement for agent configuration files follows the same principles as code quality and security enforcement. You scan where developers work, treat clarity as a quality issue, require explicit failure behavior, and apply security policy to agent configs.
Scan agent files where developers work. IDE feedback catches issues early. Git hooks provide fast local feedback. Pull request checks give review visibility. CI/CD enforcement ensures organization-wide consistency.
Treat ambiguity as a quality issue. Flag undefined terms, vague instructions, and overly complex sentences. Encourage concrete wording that an agent can follow consistently.
Require explicit failure behavior. Check for missing escape hatches. Require stop conditions for repeated failures. Require escalation paths for uncertainty or unsafe commands.
Detect contradictions and prompt bloat. Flag repeated instructions and conflicting rules. Reduce unnecessary context that increases unpredictability.
Apply security policy to agent configs. Run secrets detection on agent files. Flag dangerous command permissions. Detect PII exposure patterns. Check for exfiltration-prone instructions. Require review for permissions that allow command execution, network access, or access to sensitive paths.
Make ownership explicit. Assign maintainers for shared agent rules. Use CODEOWNERS or equivalent review routing for agent config changes. Keep organization-level policy separate from repository-specific guidance where possible.
How to harden agent config files this week
If you are looking to improve the security and reliability of your agent configuration files, here is a practical checklist:
Inventory all agent instruction files across repositories ( CLAUDE.md , .cursorrules , AGENTS.md , copilot-instructions.md , and similar)
Add agent config files to code review scope explicitly
Run linting against agent config files to catch ambiguity, missing escape hatches, and security issues
Remove hardcoded secrets , tokens , customer data , and production examples
Replace vague instructions with concrete rules
Add escape hatches for failed tests, missing context, command failures, and uncertainty
Remove duplicate or contradictory instructions
Restrict dangerous shell commands and require confirmation for destructive operations
Add PR checks so agent config changes cannot bypass review
Track recurring findings across repositories to identify shared policy problems
The path forward is straightforward. Write clearer instructions. Define failure behavior. Remove dangerous permissions and sensitive data. Enforce checks consistently across repositories.
If agent instructions affect production code, they deserve the same engineering discipline as production code.
Run AgentLinter across your connected repositories to identify ambiguity, missing escape hatches, and security risks in AI agent configuration files.
Scan your repository for free →
The Visibility Problem Behind AI Tool Adoption in Engineering Teams
Executives are asking engineering teams to accelerate AI adoption. Most engineering leaders, though, cannot answer a basic question: which AI coding...
What Auditors Will Ask About AI-Generated Code in 2026
Developers are already using AI coding tools. The question auditors are starting to ask is not whether your team has an AI policy, but whether you can...
Subscribe to our blog
Stay updated with our mo

[truncated]
