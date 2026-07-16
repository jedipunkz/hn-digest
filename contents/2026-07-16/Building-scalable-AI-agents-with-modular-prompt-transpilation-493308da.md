---
source: "https://developers.googleblog.com/building-scalable-ai-agents-with-modular-prompt-transpilation/"
hn_url: "https://news.ycombinator.com/item?id=48936621"
title: "Building scalable AI agents with modular prompt transpilation"
article_title: "Building scalable AI agents with modular prompt transpilation\n- Google Developers Blog"
author: "yruzin"
captured_at: "2026-07-16T17:04:57Z"
capture_tool: "hn-digest"
hn_id: 48936621
score: 4
comments: 1
posted_at: "2026-07-16T16:22:09Z"
tags:
  - hacker-news
  - translated
---

# Building scalable AI agents with modular prompt transpilation

- HN: [48936621](https://news.ycombinator.com/item?id=48936621)
- Source: [developers.googleblog.com](https://developers.googleblog.com/building-scalable-ai-agents-with-modular-prompt-transpilation/)
- Score: 4
- Comments: 1
- Posted: 2026-07-16T16:22:09Z

## Translation

タイトル: モジュール式プロンプト変換によるスケーラブルな AI エージェントの構築
記事のタイトル: モジュール式プロンプト変換を使用したスケーラブルな AI エージェントの構築
- Google 開発者ブログ
説明: プロンプトをビルド アーティファクトとして扱うことで、AI エージェントをスケールします。モジュール式テンプレート、トランスパイラー、CI/CD 検証を使用して実行時エラーを防ぐ方法を学びます。

記事本文:
コミュニティ/イベント
学ぶ
ブログ
YouTube
検索
コミュニティ/イベント
モジュール式プロンプトトランスパイルを備えたスケーラブルな AI エージェントの構築
初めて AI エージェントを構築するときは、通常、単一のモノリシック システム プロンプトで問題ありません。いくつかの命令 (おそらく 1 つまたは 2 つのツール定義) があり、すべてが 1 つの読み取り可能なファイル内に存在します。
しかし、本番環境で使用し始めると、その形式は単に壊れてしまいます。チームは、安全ポリシー、ドメイン固有のルール、フォーマット要件、エスカレーション動作を階層化し始めます。突然、エージェントのコントロール プレーン全体が 1 つの命令ファイル内に存在することになり、まさにそこから問題が始まります。
これは、ソフトウェア エンジニアリングの古典的なスケーリング問題です。すべての懸念事項を 1 つのファイルに押し込むと、システムについて推論する能力が失われます。コラボレーションは悪夢となり、テストは難しくなり、あるワークフローを改善するための小さな変更が、別のワークフローを静かに破壊する可能性があります。
実稼働規模では、迅速な保守性がエージェントの信頼性となります。
モノリシックプロンプトが壊れる理由
プロンプトが特定のサイズを超えると、通常、次の 3 つの主な障害モードが発生します。
曖昧な爆発範囲: 標準的なソフトウェア エンジニアリングでは、レビュー担当者がモジュールの境界、呼び出しサイト、テストを通じて変更の範囲を推論するのは簡単です。ただし、システムプロンプトの差分はさらに困難です。文を追加すると、エージェント全体に予期せぬ副作用が生じる可能性があり、多くの場合、予測やテストが困難です。
コピー＆ペーストのドリフト: 組織が拡大するにつれて、多くのチームは、内部サービスの使用手順、PII の処理、安全ポリシー、エスカレーション プロトコルなど、さまざまなアプリケーションの共有ロジックを複製することになります。これにより、同じ機能のコピーアンドペーストや複数のバージョンが発生し、不整合が発生します。
遅延ランタイム

エラー: スプロール化を管理するために、チームはアドホックな文字列フォーマットや単純なテンプレートに頼ることがよくあります。これはオーサリングに役立ちますが、エラー検出がランタイムにプッシュされます。変数の欠落または無効なインポート パスが原因で、めったに使用されない特定のワークフローがトリガーされた場合にのみ失敗するプロンプトをデプロイできます。
テンプレートは良い出発点ですが、それだけでは十分ではありません。実稼働システムには、確定的なビルド、静的検証、CI/CD 統合が必要です。
プロンプトをソフトウェア成果物のように扱う
ここでの解決策は、単なる静的テキストではなく、プロンプトをビルド成果物のように扱うことです。
1 つのモノリシック プロンプト ファイルを維持する代わりに、モジュール式のスキル ファイルを作成できます。これにより、各ファイルの範囲を縮小し、特定の動作をカプセル化できるため、チームが懸念事項を分離し、コンポーネントを個別に反復できるようになります。
最上位のエージェント プロンプト テンプレートは次のようになります。
# Agents/sre_agent.prompt.md (プロンプトテンプレートファイル)
{% 「shared/safety.prompt.md」を含む %}
{% 「shared/tool_usage.prompt.md」を含めます %}
あなたは、{{ 環境 }} 環境で活動する SRE トリアージ エージェントです。
{% ifallow_remediation %}
修復手順を推奨することはできますが、破壊的な行為には人間の承認が必要です。
{% その他 %}
問題を検査、要約、説明することはできますが、修復アクションは推奨しません。
{% endif %}
{% マクロ Bullet_section(title, items) %}
## {{ title.rstrip() }}
{アイテム内のアイテムの%}
- {{ item.rstrip() }}
{% 終了の %}
{% 終了マクロ %}
{{ Bullet_section("必要な調査手順", [
"最近の展開イベントを検査する",
"レイテンシーまたはエラー率の変化についてサービスメトリクスを確認します",
「繰り返される失敗パターンのログを確認する」
]) }}
プレーンテキスト
コピーされました
これにより、両方の長所が得られます。テンプレート層を使用すると、共有命令を作成できます。

環境固有の値を挿入し、マクロを使用します。しかし、ビルド システムの場合、すべてのインクルードは依存関係であり、すべての変数は要件です。その結果、決定論的で完全にレンダリングされたアーティファクトが得られ、モデルに到達する前にテスト、監査、および差分を行うことができます。次に、トランスパイラーを使用してテンプレートのインポートを解決し、エージェントが取り込めるファイルを生成します。
たとえば、environment =production、allow_remediation = true の場合、トランスパイルされたアーティファクトは次のようになります。
あなたは、運用環境で活動する SRE トリアージ エージェントです。
修復手順を推奨することはできますが、破壊的な行為には人間の承認が必要です。
## 必要な調査手順
- 最近の展開イベントを検査する
- サービスメトリックの遅延やエラー率の変化を確認する
- 繰り返される失敗パターンのログを確認する
プレーンテキスト
コピーされました
高レベルのトランスパイル パイプラインは次のようになります。
ビルド時の検証は必須です
実稼働グレードのトランスパイラは、実行前にエラーをキャッチする必要があります。
ビルド プロセス中に、欠落しているインポート、未定義の変数、循環依存関係の検証チェックを実行する必要があります。ここでは依存関係グラフが非常に貴重であるため、堅牢なテンプレート エンジンの必要性が強化されます。各プロンプト フラグメントを有向グラフ内のノードとして扱うと、運用環境でサイレント エラーを引き起こす再帰的なインポートを簡単に捕捉できます。
これにより、ドリフトチェックも可能になります。ソース (ゴールデン ファイルと呼ばれる) からトランスパイルされたプロンプトを再生成し、現在コミットされているアーティファクトと比較できるように CI パイプラインを設定できます。出力が異なる場合、ビルドは失敗します。これにより、リポジトリ内のコードが実稼働環境で実行されているものとまったく同じであることが保証され、ソース ファイルとデプロイメント間のギャップが排除されます。

巧妙に作られたアーティファクト。
ダイナミックなスキルとエージェント作成のアップデート
モジュール式プロンプト フラグメントのスキル ライブラリが増加するにつれて、必ずしもすべてのエージェントが毎回すべてのスキルをロードする必要はありません。これを行うとトークンが消費され、エージェントのタスク固有のパフォーマンスに干渉する可能性のあるノイズが発生します。
より良いアーキテクチャ パターンは、漸進的な開示を活用することです。ここで、安定したコントロール プレーンをタスク固有のコンテキストから分離します。コンパイルされた基本プロンプトは、アイデンティティや安全境界などの交渉不可能な動作を強制する必要があります。その後、実行時にエージェントはツールを使用して、当面のタスクに必要な特定のスキル モジュールのみを動的に取得できます。これにより、コンテキストの消耗が軽減され、エージェントがそのタスクに集中し続けることができます。
このモジュラー システムを取得すると、強力なワークフローが解放されます。エージェントは独自の命令レイヤーの維持を支援し、自立したエージェント システムの作成に役立ちます。エージェントが新しいタイプのインシデントを解決すると、理論的には新しいスキル モジュールを作成し、関連するインポートを更新し、プル リクエストを開くことができます。
エージェントは自身の命令をリアルタイムで変更しません。コードの変更を提案しています。次に、トランスパイラーは、その提案に対して、他のコード変更と同じ厳格な検証とレビューを行います。人間のレビュー担当者は PR を検査し、評価を実行し、変更をマージできます。
プロダクション プロンプト トランスパイラーは、プロンプト エンジニアリングをビルド システムの問題として再構成します。
モジュール式スキル ファイルを構築すると、標準のソフトウェア インフラストラクチャの場合と同様に、依存関係を解決し、インポートを検証し、ドリフト チェックを強制できます。変更が当社の既存の検証およびレビュープロセスを通過する場合、エージェントは独自のロジックの改善を提案できるようになります。
AI エージェントが深く統合されるにつれて

重要なワークフローに組み込まれる場合、その命令層には、私たちがソフトウェアに要求するのと同じ信頼性基準が必要です。プロンプトは編集するだけでなく、構築、検証、バージョン管理、デプロイする必要があります。
AI
クラウド
お知らせ
ベストプラクティス
Google のエージェント開発キットと A2A を使用して、言語を超えたマルチエージェント チームを構築する
AI
クラウド
お知らせ
ソリューション
Gemini Enterprise Agent プラットフォームの選択肢の拡大: 並列 Web 検索によるグラウンディングの導入
AI
クラウド
ハウツーガイド
ベストプラクティス
ADK 2.0を構築した理由
ウェブ
AI
お知らせ
学ぶ
LiteRT.js、Google の高性能 Web AI 推論
プログラム
Google 開発者プログラム
開発者コンソール
Google APIコンソール

## Original Extract

Scale your AI agents by treating prompts as build artifacts. Learn how to use modular templates, transpilers, and CI/CD validation to prevent runtime errors.

Community/Events
Learn
Blog
YouTube
Search
Community/Events
Building scalable AI agents with modular prompt transpilation
When you’re first building an AI agent, a single, monolithic system prompt is usually fine. You have a few instructions, maybe a tool definition or two, and everything lives in one readable file.
But as you start using them for production purposes, that format simply just breaks down. Teams start layering on safety policies, domain-specific rules, formatting requirements, and escalation behaviors. Suddenly, you have your entire agent’s control plane within a single instruction file which is exactly where the trouble starts.
This is a classic software engineering scaling problem. When you push every concern into a single file, you lose the ability to reason about the system. Collaboration becomes a nightmare, testing gets finicky, and a small change meant to improve one workflow can quietly break another.
At production scale, prompt maintainability becomes agent reliability.
Why monolithic prompts break down
We typically see three main failure modes when prompts grow beyond a certain size:
Obscured blast radius: In standard software engineering, it’s easy for reviewers to reason about the scope of a change through module boundaries, call sites, and tests. System prompt diffs are harder though. Adding a sentence could have unintended side effects across the entire agent, which is often hard to predict or test.
Copy-paste drift: As organizations scale, many teams end up duplicating shared logic for various applications such as internal service usage instructions, PII handling, safety policies, or escalation protocols. This leads to copy-pasting or multiple versions of the same functionality leading to inconsistencies.
Deferred runtime errors: To manage the sprawl, teams often resort to ad-hoc string formatting or simple templates. While this helps with authoring, it pushes error detection to runtime. You might deploy a prompt that only fails when a specific, rarely-used workflow is triggered because of a missing variable or an invalid import path.
Templates are a good start, but they aren't enough. Production systems require deterministic builds, static validation, and CI/CD integration.
Treat prompts like software artifacts
The solution here is to treat prompts like build artifacts as opposed to just static text.
Instead of maintaining one monolithic prompt file, you can author modular skill files. This allows you to reduce the scope of each file and encapsulate a specific behavior, which allows teams to separate concerns and iterate on components individually.
A top-level agent prompt template might look something like this:
# agents/sre_agent.prompt.md (prompt template file)
{% include "shared/safety.prompt.md" %}
{% include "shared/tool_usage.prompt.md" %}
You are an SRE triage agent operating in the {{ environment }} environment.
{% if allow_remediation %}
You may recommend remediation steps, but destructive actions require human approval.
{% else %}
You may inspect, summarize, and explain the issue, but do not recommend remediation actions.
{% endif %}
{% macro bullet_section(title, items) %}
## {{ title.rstrip() }}
{% for item in items %}
- {{ item.rstrip() }}
{% endfor %}
{% endmacro %}
{{ bullet_section("Required investigation steps", [
"Inspect recent deployment events",
"Check service metrics for latency or error-rate changes",
"Review logs for repeated failure patterns"
]) }}
Plain text
Copied
This gives you the best of both worlds. The templating layer lets you compose shared instructions, inject environment-specific values, and make use of macros. But for the build system, every include is a dependency, and every variable is a requirement. The result is a deterministic, fully rendered artifact that you can test, audit, and diff before it ever reaches the model. We can then use a transpiler to resolve the template imports to generate a file that is ready to be ingested by an agent.
For example, if environment = production and allow_remediation = true, the transpiled artifact would look like this:
You are an SRE triage agent operating in the production environment.
You may recommend remediation steps, but destructive actions require human approval.
## Required investigation steps
- Inspect recent deployment events
- Check service metrics for latency or error-rate changes
- Review logs for repeated failure patterns
Plain text
Copied
A high-level transpilation pipeline would look something like this:
Build-time validation is mandatory
A production-grade transpiler should catch errors before runtime.
We should be running validation checks for missing imports, undefined variables, and circular dependencies during the build process. Dependency graphs are invaluable here, reinforcing the need for a solid template engine. If you treat each prompt fragment as a node in a directed graph, you can easily catch recursive imports that would otherwise cause a silent failure in production.
This also enables drift checking. You can set your CI pipelines to be able to regenerate the transpiled prompt from source (referred to as the golden file) and compare it against the currently committed artifact. If the outputs differ, the build fails. This ensures that the code in your repo is exactly what’s running in production, eliminating the gap between source files and deployed artifacts.
Dynamic skills and agent-authored updates
As your skill library of modular prompt fragments grows, you don't necessarily want every agent to load every skill every time. Doing so consumes tokens and introduces noise that can interfere with the agent's task-specific performance.
A better architectural pattern would be to leverage progressive disclosure. This is where we separate the stable control plane from task-specific context. The compiled base prompt should enforce non-negotiable behaviors like identity and safety boundaries. Then, at runtime, the agent can use a tool to dynamically retrieve only the specific skill modules required for the task at hand; this reduces context exhaustion and helps to keep the agent focused on its task.
Once you have this modular system, you unlock a powerful workflow: agents can help maintain their own instruction layer, helping to create a self-sustaining agentic system. When an agent resolves a new type of incident, it could theoretically draft a new skill module, update the relevant imports, and open a pull request.
The agent isn't mutating its own instructions in real-time; it's proposing a code change. The transpiler then subjects that proposal to the same validation and review rigors as any other code change. A human reviewer can inspect the PR, run the evals, and merge the change.
A production prompt transpiler reframes prompt engineering as a build-system problem.
When we build modular skill files, we can resolve dependencies, validate imports, and enforce drift checks just as we do with our standard software infrastructure. Agents become capable of suggesting improvements to their own logic, provided those changes pass through our existing validation and review processes.
As AI agents become deeply integrated into critical workflows, their instruction layers need the same reliability standards we demand of our software. Prompts shouldn't just be edited, they should be built, validated, versioned, and deployed.
AI
Cloud
Announcements
Best Practices
Build Cross-Language Multi-Agent Team with Google’s Agent Development Kit and A2A
AI
Cloud
Announcements
Solutions
Expanding Choice in Gemini Enterprise Agent Platform: Introducing Grounding with Parallel Web Search
AI
Cloud
How-To Guides
Best Practices
Why we built ADK 2.0
Web
AI
Announcements
Learn
LiteRT.js, Google's high performance Web AI Inference
Programs
Google Developer Program
Developer consoles
Google API Console
