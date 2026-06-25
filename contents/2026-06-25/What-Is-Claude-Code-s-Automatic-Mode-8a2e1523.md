---
source: "https://www.polimetro.com/en/What-is-Claude-Code%27s-automatic-mode/"
hn_url: "https://news.ycombinator.com/item?id=48671975"
title: "What Is Claude Code's Automatic Mode"
article_title: "What is Claude Code's automatic mode and when should you use it?"
author: "Gedxx"
captured_at: "2026-06-25T11:53:09Z"
capture_tool: "hn-digest"
hn_id: 48671975
score: 3
comments: 0
posted_at: "2026-06-25T11:39:08Z"
tags:
  - hacker-news
  - translated
---

# What Is Claude Code's Automatic Mode

- HN: [48671975](https://news.ycombinator.com/item?id=48671975)
- Source: [www.polimetro.com](https://www.polimetro.com/en/What-is-Claude-Code%27s-automatic-mode/)
- Score: 3
- Comments: 0
- Posted: 2026-06-25T11:39:08Z

## Translation

タイトル: クロード・コードの自動モードとは
記事のタイトル: クロード コードの自動モードとは何ですか?いつ使用する必要がありますか?
説明: クロードコードをどのようにコード化するかを詳しく説明します。

記事本文:
クロード コードの自動モードとは何ですか?いつ使用する必要がありますか?
コンテンツにスキップ
マルチメーター
エンジニアリング ソフトウェアとツール
メイン » 人工知能 » クロード コードの自動モードとは: AI 支援開発における仕組み、利点、および制限
Claude Code の自動モードとは何ですか: AI 支援開発における仕組み、利点、制限事項
Claude Code の自動モードを使用すると、AI が権限について決定し、開発タスクの自律性とセキュリティのバランスを取ることができます。
これには、潜在的に危険なアクションをブロックし、安全なコマンドのみを通過させる AI 分類子が組み込まれているため、手動による承認の必要性が制限されます。
適切な使用方法は、タスクの種類、環境、エージェントの信頼レベルによって異なります。信頼されたリポジトリ内で反復的なタスクを行う場合には、この使用を強くお勧めします。
現在、チーム プラン ユーザー向けの試用モードで利用できますが、エンタープライズおよび API に拡張する予定です。ただし、リスクを完全に排除するわけではなく、監督が必要です。
最近、AI 支援ソフトウェア開発におけるインテリジェントな自動化は、Claude Code の自動モードの開始により大幅に前進しました。技術界や開発者コミュニティで大きな反響を呼んだこの機能は、コード プロジェクトにおける AI の自律性と運用上のセキュリティのバランスをとるための新しいアプローチを導入し、固有のリスクに留意しながら反復的なタスクを高速化します。
過去数か月にわたって、さまざまな専門メディアやブログがこの機能を徹底的に分析し、その長所、限界、推奨される使用例を強調してきました。 Anthropic It の主力製品である Claude Code がリリースされ、スタートアップ チームや開発者の間で野火のように広がり、テッド サービスの負荷を軽減するために設計された新機能がリリースされました。

危険な行為をブロックし、安全な行為のみを許可する高度な分類システムのおかげで、常に許可を付与し、破壊的なコマンドによって引き起こされる災害を防ぐことができます。しかし…自動モードとは正確には何で、どのように機能し、どのようなシナリオで有効にする価値があるのでしょうか?
クロード コードの自動モードは何で構成されていますか?
Claude Code の自動モードでは、各ステップで人間の介入を必要とせずに、AI 自体が特定のアクションをいつ実行するかを決定できます。このアプローチは、人工知能がすべてを操作するための完全な決定権を持っていることを意味するものではなく、むしろ AI 分類器を使用した自動化されたリスク分析の層を導入します。このようにして、安全とみなされるタスクは自律的に実行されますが、ディレクトリの削除、重要なファイルの移動、機密データの流出の可能性など、損害を伴う可能性のあるタスクは自動的にブロックされるか、ユーザーに対して新しい許可リクエストが生成されます。
ここでの目標は、プログラミングと自動化のワークフローを最適化し、開発者がすべての些細な操作を手動で承認する必要をなくすと同時に、突破できない一線としてセキュリティを維持することです。有名な危険なスキップ権限などのオープン機能 これらはクロード コードにすでに存在していましたが、自動モードはより中間的で賢明なソリューションであり、生産性を犠牲にすることなく制御を求める人にとって理想的です。
技術的な運用と許可に関するポリシー
自動モード エンジンは、Claude Code の AI に統合されたインテリジェントな分類器に基づいています。シェル、ファイル編集、リポジトリとの対話など、ツールを呼び出すたびに、システムは最初に潜在的な影響を評価します。リスク基準が許容範囲内にある場合 (たとえば、t 以内の小さな変更など)

グローバル環境を変更しない同じブランチまたはリファクタリング)、アクションは承認され、特定のユーザーの承認を必要とせずに直接実行されます。 。
逆に、分類子が潜在的な危険の兆候 (大量削除の試み、悪意のあるコードの実行、割り当てられたディレクトリ外への移動、システムの整合性を危険にさらすアクティビティなど) を識別した場合、操作は自動的に停止されます。このような場合、AI はより安全な代替パスを試みるか、問題が解決せずにいくつかのブロックに到達した場合は、最終的に新しい手動許可リクエストを発行し、制御と決定をプログラマに戻します。
この機能がスタートアップ企業、プログラマー、技術チームにとって重要なのはなぜですか?
AI 支援開発、特に反復的なプロセスの自動化は、生産性の大幅な飛躍を表します。しかし、この革命に対する主な障害の 1 つは、AI が暴走してプロジェクトにとって致命的なコマンドを実行する可能性があるという懸念から、セッションごとに数十のアクションを承認する必要があることです。自動モードは、この従来の緊張を解決するために正確に設計されています。
AI がリポジトリ内のローカル タスク (リファクタリング、インポートの更新、依存関係の管理、テストの修正など) のみを実行する必要がある長期プロジェクトの摩擦を軽減し、小さなステップごとの権限の煩わしさを排除します。
危険な代替手段は避けてください。 AI はデフォルトで本当に危険なものをフィルタリングしてブロックするため、すべての権限のバイパス (大惨事につながる可能性があります) を避けるためです。
テストと修正のサイクルの流動性が向上し、開発者をアクティブな立場に保ちながら、些細な監督から解放されます。
価値があるケース: 自動モードに最適なタスク
自動モードは、ボトルが

ネックとなるのは承認の負担であり、タスク自体に伴うリスクではありません。たとえば、次のような場合に最適です。
機械的だが関連性のある作業に AI が依存している、監査済みのプロジェクト内での一貫した広範なリファクタリング (抽象化の名前変更、インポートの更新、ブランチ全体のクリーンアップ)。
設備のメンテナンス - 特にロックファイルが明確に定義されており、同じ環境内に誤って操作される可能性のある重要なスクリプトが混在していない場合。
AI が小さな修正を適用してテストを繰り返し実行する、テストと修正の古典的なループが自動化されます。
AI が安全で変更可能なスペース (テスト用の分離された機能ブランチまたはフォークなど) に限定される、長く反復的なセッション。
ユーザー エクスペリエンスと Anthropic の公式ドキュメントによると、重要なのは、タスクが明確に定義されており、簡単に確認できることです。したがって、エラーによる潜在的な損害は、現在のリポジトリまたはブランチの範囲に限定され、重要なシステム操作に波及したり、実際の運用環境に影響を与えたりすることはありません。
自動モードを有効にしてはいけない場合: リスクと悪い習慣
特定のシナリオでは、自動モードがその有用性をすべて失い、慎重でないチームにとっては問題の原因になる可能性さえあります。一般に、次のような場合には使用しないでください。
運用、共有インフラストラクチャ、重要なデータベース、またはローカル プロジェクト外の機密リソースを変更できるコマンドに関連するタスク。
デプロイメント スクリプト、認証情報ファイル、グローバル構成ファイル、または単一の間違った手順によってエラーが伝播されるマルチサービス統合を含む、大まかに定義されたリポジトリ。
最高の権限を持つ上級オペレーターのみが環境を操作できるすべての状況 (たとえば、パーマに影響を与える変更など)

問題、データベース移行、または実稼働環境への自動デプロイメントなど）。
Claude Code は、その目的に最適な acceptEdits などの侵入性の低いモードを提供するため、単純な編集の承認を減らすことだけが目的の場合。
どのレベルのセキュリティと制御が提供されますか? AIに権限を委譲するのは信頼できるのか？
Anthropic は「責任ある自律性」の哲学に特に重点を置き、強力ではあるものの確実ではないアルゴリズムによる監視の層を追加しました。権限分類子は、他の AI ベースのシステムと同様に、コンテキスト エラーを引き起こす可能性があります。
システムがエージェントのリクエストを誤って解釈すると、本当に危険なアクションが発生する可能性があります。
正当なトランザクションが散発的にブロックされ (誤検知)、手動ワークフローへの復帰を余儀なくされる可能性があります。
このため、現時点で推奨されている使用方法は、実際の運用シナリオに実装するずっと前に、テスト環境 (「サンドボックス」) で使用することになります。貴重なスクリプト、重要なコードベース、または回復が困難なリソースを所有する企業は、AI に引き継がせる前にバックアップ、スナップショット、または ZIP バックアップ ファイルを作成することが多く、これにより予期せぬリスクに対する回復力が高まります。
この機能は開発者の技術的判断に代わるものではなく、監視をより高いレベルに移行します。どのような条件を信頼するか、いつ手動介入を要求するかをユーザーが決定するため、自動化をセキュリティの絶対的な保証として扱うべきではありません。
利用可能なモード、要件、および特定の操作
現在、自動モードはリサーチ プレビュー段階にあり、チーム プラン ユーザーに対してのみ有効です。このオプションが利用できない場合は、適切なプランを持っていない、管理者がまだ有効にしていない、互換性のないモデルを使用している (Sonnet 4.6 および Opus 4.6 で動作します)、または環境が整っている可能性があります。

公式の Claude Code CLI またはクラウド セッション以外のセッション。
この機能は間もなくエンタープライズとその他の顧客の両方に拡張される予定で、同社自身もドキュメントの中で、ローカル バージョン、クラウド、またはさまざまな統合開発環境 (CLI、VS Code、リモート コントロールなど) によって制限、インターフェイス、または許可モードが異なる可能性があると述べています。
特に、自動モードをアクティブにすると、トークンの消費と待ち時間の両方で追加のコストがかかる可能性があります。 (つまり、分類レイヤーが各コマンドに追加のチェックを追加するため、応答時間は多少長くなります。) したがって、この機能を無料のショートカットとしてではなく、ワークフローの質的改善として考慮し、非常に長いセッション中にその影響を確認することをお勧めします。
自動モードがいつ、誰に推奨されるか
次の条件が満たされる場合、自動モードをアクティブにすることが有利です。
この機能への実際のアクセスが可能です (チーム プラン、または近日中にエンタープライズ/API)。
このタスクは非常に長いため、手動で承認すると大幅な時間のロスが発生します。
作業環境は明確に定義されており、エラーは現在のブランチまたはリポジトリに限定されるという高い信頼性があります。
最終結果もエンジニアによるレビューを受けて承認されるため、やみくもにすべての権限がAIに委任されるわけではない。
ただし、次のような場合には使用をお勧めしません。
サブスクリプション プランで許可されていないか、まだサポートされていない個人環境で動作しています。
このタスクは、本番環境、共有インフラストラクチャ、またはリポジトリ構造によって定義された「信頼境界」の外側にあるファイルに影響します。
他のマイナー モードの直接の代替として使用することを目的としているのでしょうか、それとも単にローカル編集を高速化するためのものでしょうか (acceptEdits で十分な場合があります)。
現在の障壁と課題

自動化された方法の
自動モードには利点がありますが、リスク要素が排除されるわけでも、完全な制御が保証されるわけでもありません。以前のバージョンを試したユーザーや専門家も指摘しているように、この関数は依然として機密コンテキストを識別する分類子の強度に依存しています。
環境の範囲が割り当てられたフォルダーのみよりも広い場合、クロード コードによってコードベースが破損することがあります。
システムは、過度の熱意 (安全なタスクのブロック) と不注意 (問題のあるコマンドの実行) の両方によって間違いを犯す可能性があります。
自律性が高まっている一方で、慎重な監査、監視、レビューの必要性も高まっていることを理解する必要があります。エージェントがどのように機能したか。以前は無限の手動認証を必要としていた長時間のセッションがよりスムーズに実行できるようになりましたが、重要なプロジェクトでは二次的な制御や適切なバックアップ戦略なしにやみくもにアクティブ化すべきではありません。
他のコードエージェントとの比較とAI分野の動向
Claude Code での自動モードの登場は、コード アシスタント市場の一般的な傾向に対応しています。 GPT-4 コード インタプリタやその他の自律エージェントなどの高度なツールは、システムの限界についての懸念と議論を引き起こしています。

[切り捨てられた]

## Original Extract

Discover in detail how Claude Code

What is Claude Code's automatic mode and when should you use it?
Skip to content
Multimeter
Engineering Software and Tools
Main » Artificial Intelligence » What is Claude Code's Automatic Mode: how it works, its advantages, and its limitations in AI-assisted development
What is Claude Code's Automatic Mode: how it works, its advantages, and its limitations in AI-assisted development
Claude Code's Automatic Mode allows AI to make decisions about permissions, balancing autonomy and security in development tasks.
It incorporates an AI classifier that blocks potentially dangerous actions and only allows safe commands to pass through, limiting the need for manual approvals.
Appropriate use depends on the type of task, environment, and level of trust in the agent, being highly recommended for repetitive tasks within trusted repositories.
Currently available in trial mode for Team plan users, with plans to expand to Enterprise and API However, it does not completely eliminate the risk and requires supervision.
Recently, intelligent automation in AI-assisted software development has taken a significant leap forward with the launch of Claude Code's Automated Mode. This feature, which has generated considerable excitement in tech circles and developer communities, introduces a novel approach to balancing AI autonomy with operational security in code projects, accelerating repetitive tasks while remaining mindful of inherent risks.
Over the past few months, various specialized media outlets and blogs have thoroughly analyzed this feature, highlighting its strengths, limitations, and recommended use cases. Claude Code, a flagship product of anthropic It has launched, spreading like wildfire among startup teams and developers, a new feature designed to reduce the tedium of constantly granting permissions and prevent disasters caused by destructive commands, thanks to an advanced classification system that blocks dangerous actions and only allows safe ones. But… what exactly is automatic mode, how does it work, and in what scenarios is it worth activating?
What does Claude Code's Automatic Mode consist of?
Claude Code's automatic mode allows the AI ​​itself to decide when to perform certain actions without requiring human intervention at each step. This approach does not imply that artificial intelligence has carte blanche to manipulate everything, but rather introduces a layer of automated risk analysis using an AI classifier. In this way, tasks considered safe They are performed autonomously, while those that may involve damage — such as the deletion of directories, critical file movements or the possible exfiltration of sensitive data — are automatically blocked or generate a new permission request for the user.
The goal here is to optimize workflows in programming and automation, relieving developers of having to manually approve every trivial operation, while maintaining security as an unbreakable red line. Open functionalities like the famous dangerously skip permissions They already existed in Claude Code, but automatic mode represents a much more intermediate and sensible solution, ideal for those seeking control without sacrificing productivity.
Technical operation and permit policy
The automatic mode engine is based on an intelligent classifier integrated into Claude Code's AI. In every call to a tool—whether in the shell, file editing, or interaction with repositories—the system first assesses the potential impact. If the risk criteria are within an acceptable range (for example, small modifications within the same branch or refactorings that do not alter the global environment), The action is approved and executed directly without requiring specific user approval. .
Conversely, if the classifier identifies signs of potential danger (such as mass deletion attempts, execution of malicious code, movements outside the assigned directory, or activities that jeopardize the integrity of the system), the operation is automatically stopped. In such cases, the AI ​​can attempt a safer alternative path or, if the problem persists and several blocks are reached, it finally raises a new manual permission request, returning control and the decision to the programmer.
Why is this feature relevant for startups, programmers, and technical teams?
AI-assisted development, and especially the automation of repetitive processes, represents a huge leap in productivity. However, one of the main obstacles to this revolution has been the requirement to approve dozens of actions per session, for fear that the AI ​​could run wild and execute commands that could be fatal to the project. The automatic mode is precisely designed to resolve this traditional tension:
reduces friction in long-term projects where AI only needs to perform local tasks within the repository (refactoring, updating imports, managing dependencies, correcting tests, etc.), eliminating the bureaucracy of permissions for each small step.
Avoid the dangerous alternative. to avoid bypassing all permissions (which could end in disaster), since AI filters and blocks what is truly risky by default.
Improves fluidity of test and fix cycles , keeping the developer in an active position, but relieving them of trivial supervision.
Cases where it is worthwhile: tasks ideal for Automatic Mode
Automatic mode makes sense when the bottleneck is the burden of approvals, not the risk associated with the task itself. For example, it's perfect for:
Consistent and extensive refactors (renaming abstractions, updating imports, cleaning up entire branches) within already audited projects where AI is relied upon for that mechanical but relevant work.
Maintenance of facilities — especially when the lockfile is well defined and there are no critical scripts mixed in the same environment that could be manipulated by mistake.
Classic loops of test/fix automated, where AI applies small corrections and runs tests over and over again.
Long and repetitive sessions in which the AI ​​is limited to a safe and revisable space (e.g., an isolated feature branch or forks for testing).
The key, according to user experience and Anthropic's official documentation, is that The tasks are well-defined and easily reviewable The potential damage from an error is therefore confined to the scope of the current repository or branch, without being able to spread to critical system operations or affect live production environments.
Where NOT to activate Auto Mode: risks and bad practices
There are certain scenarios where automatic mode loses all its usefulness and can even be a source of problems for less cautious teams. In general, it should not be used in:
Tasks related to production, shared infrastructure, critical databases, or commands capable of modifying sensitive resources outside of the local project.
Loosely defined repositories containing deployment scripts, credential files, global configuration files, or multi-service integrations where a single wrong step propagates the error.
All those situations where only a senior operator with maximum authority would manipulate the environment (for example, changes that affect permissions, database migrations, or automatic deployments to production).
Cases where the only interest is reducing simple edit approvals, since Claude Code offers less intrusive modes, such as acceptEdits , ideal for that purpose.
What level of security and control does it offer? Is it reliable to delegate authority to AI?
Anthropic has placed special emphasis on the philosophy of 'responsible autonomy', adding a layer of algorithmic supervision that, while powerful, is not infallible. The permission classifier, like any other AI-based system, can make context errors:
Some truly dangerous actions could end up happening if the system misinterprets the agent's request.
Legitimate transactions could be blocked sporadically (false positives), forcing a return to a manual workflow.
For this reason, The recommended use for now remains in test environments ('sandbox') long before implementing it in real-world production scenarios. Those with valuable scripts, critical codebases, or hard-to-recover resources often create backups, snapshots, or ZIP backup files before letting AI take over, thus increasing resilience to unexpected risks.
The function does not replace the developer's technical judgment, but rather shifts supervision to a higher level. The user still decides under what conditions to trust and when to demand manual intervention, so automation should never be treated as an absolute guarantee of security.
Available modes, requirements and specific operation
Currently, automatic mode is in research preview and is only enabled for Team plan users. If you don't see this option available, it's likely that you don't have the right plan, the administrator hasn't activated it yet, you're using an incompatible model (it works with Sonnet 4.6 and Opus 4.6), or you're in an environment other than the official Claude Code CLI or cloud sessions.
The feature is expected to be extended soon to both Enterprise and other customers, and the company itself notes in its documentation that restrictions, interface or permission modes may vary between local versions, clouds or different integrated development environments (CLI, VS Code, Remote Control…).
Notably Activating automatic mode may involve additional costs in both token consumption and latency. (That is, the response time will be somewhat longer, since the classification layer adds an extra check to each command.) Therefore, it's advisable to consider this feature not as a free shortcut, but as a qualitative improvement to the workflow, and to review its impact during very long sessions.
When and for whom is automatic mode recommended
Activating automatic mode is advantageous when these conditions are met:
Real access to the feature is available (Team plan or, soon, Enterprise/API).
The task is lengthy enough that manual approvals would result in a significant loss of time.
The working environment is well defined and there is a high level of confidence that any errors will be confined to the current branch or repository.
The final result will also be reviewed by an engineer before being accepted, and all authority is not blindly delegated to AI.
However, its use is discouraged in cases where:
The subscription plan does not allow it, or it operates in personal environments that are not yet supported.
The tasks affect production, shared infrastructure, or files outside the "trust boundary" defined by the repo structure.
Is it intended to be used as a direct replacement for other minor modes, or is it simply to speed up local editing (where acceptEdits may be enough).
Current barriers and challenges of the automated method
Despite its advantages, automatic mode does not eliminate the risk component nor ensure perfect control. As users and experts who have experimented with previous versions have also pointed out, the function still depends on the strength of the classifier to discern sensitive contexts:
Occasionally, Claude Code has damaged codebases when the scope of the environment was broader than just the assigned folder.
The system can make mistakes both by being overzealous (blocking safe tasks) and by being negligent (letting through some problematic command).
It is necessary to understand that Autonomy is growing, but so is the need for careful auditing, monitoring, and review. How the agent has worked. Long sessions, which previously resulted in endless manual authorizations, now flow better, but should never be activated blindly in critical projects without secondary controls and good backup strategies.
Comparison with other code agents and trends in the AI ​​sector
The appearance of automatic mode in Claude Code responds to the general trend within the code assistant market. Advanced tools, such as the GPT-4 Code Interpreter and other autonomous agents, are generating concern and debate about the limits of a

[truncated]
