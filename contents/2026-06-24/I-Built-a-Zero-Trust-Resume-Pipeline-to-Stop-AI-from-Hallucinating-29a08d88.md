---
source: "https://github.com/Gotili/EigenCV"
hn_url: "https://news.ycombinator.com/item?id=48666185"
title: "I Built a Zero-Trust Resume Pipeline to Stop AI from Hallucinating"
article_title: "GitHub - Gotili/EigenCV · GitHub"
author: "g0tili"
captured_at: "2026-06-24T22:27:26Z"
capture_tool: "hn-digest"
hn_id: 48666185
score: 1
comments: 0
posted_at: "2026-06-24T22:07:46Z"
tags:
  - hacker-news
  - translated
---

# I Built a Zero-Trust Resume Pipeline to Stop AI from Hallucinating

- HN: [48666185](https://news.ycombinator.com/item?id=48666185)
- Source: [github.com](https://github.com/Gotili/EigenCV)
- Score: 1
- Comments: 0
- Posted: 2026-06-24T22:07:46Z

## Translation

タイトル: AI の幻覚を阻止するゼロトラスト再開パイプラインを構築しました
記事タイトル: GitHub - Gotili/EigenCV · GitHub
説明: GitHub でアカウントを作成して、Gotili/EigenCV の開発に貢献します。

記事本文:
GitHub - Gotili/EigenCV · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ゴティリ
/
アイゲンCV
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
44℃

省略 44 コミット .devcontainer .devcontainer cv cv docs docs scripts scripts testing tools tools .cursorrules .cursorrules .gitignore .gitignore AI_START_HERE.md AI_START_HERE.md ライセンス ライセンス README.md README.md USER_GUIDE.md USER_GUIDE.md application_tracking.template.md application_tracking.template.md chatgpt_run.py chatgpt_run.py check_ats_score.py check_ats_score.py要件.txt要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
🛡️EigenCV: ゼロトラストエージェント再開パイプライン
あなたが持っていないスキルを ChatGPT に幻覚させるのはやめてください。
完全性を犠牲にすることなく、ATS に最適で高度にカスタマイズされた LaTeX 履歴書を生成するためのプロダクション グレードの Infrastructure-as-Code (IaC) パイプライン。
🤯 商用 AI ビルダー vs. EigenCV
業界標準 (商用 AI ビルダー):
AI に「この仕事に合わせて履歴書を最適化してください」と指示します。 AI はあなたの履歴書を創造的な文章の練習として扱います。それは静かにスキルを幻覚させ、役職を誇張し、エンジニアリングの成果を一般的な HR の流行語に言い換えます。結果として、ATS には勝った PDF が作成されましたが、嘘だらけのため技術面接には不合格でした。
EigenCV アプローチ (ゼロトラスト):
私たちはあなたの職歴を不変のデータベースとして扱います。 AI は厳密にはオーケストレーション層です。履歴書を書くわけではありません。データベースにクエリを実行して、最も関連性の高い、事前に検証された箇条書きを抽出します。
AI が不正行為を行い、欠けているスキルをプロファイルに幻覚させて ATS スコアを人為的に上げようとした場合、Python コンパイラーの嘘発見器がそれを阻止し、ビルドをハードクラッシュさせます。明らかな嘘が PDF に組み込まれることはありません。
🎨 トータルなカスタマイズとコントロール
上記のサンプルは 1 つの構成にすぎません。 EigenCV では完全に制御できます。
動的セクション: すべてのセクションが必須というわけではありません。 eを選択してください

カンマ区切りの cvorder 変数を編集するだけで、何を含めるかを正確に指定し、即座に並べ替えることができます (たとえば、「Education」を一番上に移動します)。
レイアウトの変更: プロフェッショナルな LaTeX テンプレート (例: Awesome-CV 、EigenCV-Modern ) をシームレスに切り替えます。
自分自身をブランド化: 応募先の企業に合わせて、カスタムの企業アクセント カラーを挿入します。- 多言語: 組み込みの翻訳マトリックスを使用して、英語、ドイツ語、またはその他の言語でアプリケーションを生成します。
🚀 EigenCV の使用方法: パスを選択してください
パス 1: 「ノーコード」ライフハック (非コード者向け / ChatGPT Plus)
EigenCV を使用するために、Python、LaTeX、または Git の知識は必要ありません。パイプライン全体を ChatGPT で直接実行できます。
このリポジトリ全体を ZIP ファイルとしてダウンロードします。
マスター データベースを構築します。ZIP をすべての古い履歴書とともに ChatGPT (高度なデータ分析が必要) または Claude にアップロードします。 AI に「これらの履歴書から私のキャリアに関する事実をすべて抽出し、EigenCV JSON データベースに入力してください」と指示します。
ジョブに応募する: ジョブの説明を貼り付け、AI に「 docs/AI_CLOUD_PROMPT.md にある指示に従ってこのジョブに応募してください。」と伝えます。
PDF をダウンロードします。ChatGPT はデータベースをジョブに自動的に照合し、JSON を生成し、chatgpt_run.py ラッパー スクリプトを実行します。専用の「クラウドセーフ」LaTeX テンプレートが含まれているため、ChatGPT はサンドボックス内で PDF を直接レンダリングし、ダウンロード リンクを提供します。
(フォールバック: ChatGPT がタイムアウトしても、.tex コードは生成されます。そのコードを無料の Overleaf アカウントにドラッグ アンド ドロップして、インスタント レンダリングを行うことができます (または、 pdflatex を使用して PDF をローカルでレンダリングします)。
パス 2: Agentic 開発者ルート (CLI および IDE)
ターミナル ワークフローを好む場合、最大限のビルド速度が必要な場合、または厳密なデータ プライバシーが必要な場合は、パイプラインをローカルで実行します。
前提条件: Python 3.11 以降と L が必要です。

aTeX ディストリビューション (例: TeX Live または MiKTeX)。あるいは、このリポジトリを VS Code で開き、[コンテナで再開く] をクリックするだけで、事前に構築された Docker 環境を使用できます。
依存関係のインストール: pip install -rrequirements.txt を実行します。
エージェントのセットアップ:
速度を最大化するには: 標準クラウド API に接続された Agentic CLI (Antigravity / Claude Code など) または Agentic IDE (Cursor / Windsurf) を使用してリポジトリを開きます。
100 % のハードコア プライバシーを実現するには: Agentic IDE をローカル モデル (Ollama や LM Studio など) に向けて、データがマシンから決して流出しないようにします。
DB を構築します。エージェントに次のように伝えます。「古い履歴書を移行してください。 AI_START_HERE.md に従ってください。」ゼロトラスト データベースを構築します。
応募: 求人説明を貼り付けて、「この求人に応募してください。 AI_START_HERE.md に従ってください。」と言います。
自動化: エージェントは自動的にプロンプ​​トをルーティングし、厳密な JSON を生成し、Python スクリプトをローカルで実行して PDF をレンダリングし、ATS スコアを計算します。
ほとんどの AI ツールは、「生成して祈る」アプローチを使用します。 EigenCV は Agentic Determinism を使用します。完璧で幻覚のないアプリケーションを保証する方法は次のとおりです。
不変マスター データベース: 履歴書をチャット ウィンドウに貼り付ける必要はありません。あなたの職歴は、構造化された JSON/Markdown データベースとしてハード ドライブ上にオフラインで保存されます。すべての実績、プロジェクト、スキルには数学的に検証可能な ID (例: proj_aws_migration ) があります。
AI オーケストレーション (脳): 仕事に応募すると、AI が職務内容を読み取り、戦略的なオーケストレーターとして機能します。テキストを書き込む代わりに、データベースに対して正確なクエリを実行し、ATS 一致スコアを最大化する ID のみを返します。検証されたプロフィール データのみに基づいて、非常に本物のカバーレターを作成します。
Deterministic Compilation (The Muscle): AI はレンダリング プロセスから完全にロックアウトされます。 Th

eeigenCV Python コンパイラーは、承認された ID リストを取得し、オフライン データベースから正確なテキストを取得し、それを素晴らしい ATS 最適化 LaTeX テンプレートに決定的に挿入します。完璧なレイアウト、幻覚なし。
シーケンス図
自動番号付け
俳優U 👤ユーザー役
🧠エージェントとしての参加者LLM
🗄️ データベースとしての参加者 DB
🐍 Python ツールとしての参加者 Py
四角形 rgb(40, 40, 60)
U、DB に関するメモ: フェーズ 0: セットアップとプライバシー
U->>Py: scrub_data.py を実行します。
Py->>DB: プライベートデータを消去
U->>LLM: 古い履歴書をアップロードしてください
LLM->>DB: JSON データベースの構築
終わり
四角形 rgb(20, 40, 20)
U,Py に関するメモ: フェーズ 1: ルーティング
U->>LLM: 職務内容を入力してください
LLM->>DB: クエリデータベース
DB-->>LLM: 有効な ID を返す
LLM->>LLM: build_config.json を生成する
終わり
長方形 rgb(60, 40, 20)
LLM,Py に関するメモ: フェーズ 2: コンパイル
LLM->>Py: cv_compiler.py を実行します。
Py,DB に関するメモ: 🏥 ヒーラーは ID を自動修正します<br>🛡️ 嘘発見器は幻覚を防ぎます
Py->>DB: 検証済みテキストを取得する
Py-->>LLM: Jinja2 経由で PDF をレンダリング
終わり
四角形 rgb(40, 60, 40)
LLM,Py に関するメモ: フェーズ 3: 検証
LLM->>Py: check_ats_score.py を実行
Py->>Py: コンパイルされた PDF を解析する
Py-->>LLM: ATS スコアを返す
LLM-->>U: 最終的な PDF とスコアを提示してください。
終わり
読み込み中
🚀 「嘘発見器」が作動中
+---------------------------- EigenCV コンパイラ ------------------------+
| build_config.json から CV をコンパイルしています... |
+----------------------------------------------------------------------------+
レイアウト テンプレートの使用: eigencv_resume.tex.j2、ロケール: en
+--------------------------- ゼロトラスト違反 ------------------------+
|ゼロトラスト違反: 不足しているスキルとして「Rust」を宣言しましたが、それは |
| CV 出力に幻覚が現れました! |
|自分が持っていないスキルを自由記述フィールドに人為的に注入することはできません。|
+------------------------------------------------

------------------------+
ValueError: 幻覚スキルが検出されました: Rust
ユーザーは幻覚スキルを削除し、再コンパイルします。
+---------------------------- EigenCV コンパイラ ------------------------+
| build_config.json から CV をコンパイルしています... |
+----------------------------------------------------------------------------+
CV が CV-Applicant-Google.tex に正常にコンパイルされました
pdflatex を使用して PDF を自動コンパイルしています...
CV-Applicant-Google.pdf が正常にコンパイルされました
ATS マッチスコア: 85.0 %
+--------------------------------------------------------------------------------+
|カテゴリー |スキル |
|-----------------+-------------------------------------------------|
|欠品 (1) |さび |
+--------------------------------------------------------------------------------+
[!] ATS ペナルティ適用: 1 つの重大なギャップが特定されました。
✨ コア USP
🛡️ ゼロトラストと EigenTruth エンジン: あなたの職歴は静的な JSON データベースに保存されます。 LLM が、ATS スコアを人為的に高めるために、ユーザーが持っていないスキルをプロフィールに幻覚させようとした場合、コンパイラーの EigenTruth エンジン (嘘発見器) がそれをキャッチし、ビルドをハードクラッシュさせます。
🔒 不変データベース: あなたの箇条書きとスキルは厳密に IMMUTABLE です。これらを自分で保守することも、LLM を使用して準備することもできますが、EigenCV パイプライン内では、AI はそれらを選択することのみが許可されており、書き換えることはできません。
✍️ 本物のカバーレター: AI はあなたのpersonal_dossier.md を使用して、あなたの実際のソフトスキルと趣味のみに基づいて超本物のカバーレターを作成し、企業の偽りの要素を排除します。
🎨 企業の自動カラーリング: AI は対象企業のコーポレート アイデンティティを自動的に推定し、一致するアクセント カラーを LaTeX 出力に動的に挿入します (または手動でオーバーライドすることもできます)。
🧮 高度な ATS エンジンとリアリティ チェック: コンパイル後の P

ython パーサーは、数学的に正確な ATS キーワード一致スコアを計算します。一方、AI エージェントは無慈悲なフィルターとして機能し、検証されたスキルに厳密に基づいて面接/オファーの確率と給与範囲を推定し、現実的な確率マトリックスを作成します。
🏥 自己修復 ID: LLM がデータベースから ID を選択するときに軽微なタイプミスをした場合 (例: aws_migration ではなく aws_mig )、コンパイラーの組み込み Rapidfuzz ヒューリスティックによって ID が自動修復され、脆弱なビルドのクラッシュが防止されます。
📄 自動 LaTeX コンパイル: LaTeX の解析が壊れたり、括弧が欠落したりすることはもうありません。 AI は厳密に型指定された Pydantic JSON スキーマを生成し、決定論的に美しい Jinja2 LaTeX テンプレートにコンパイルされます。
🌍 多言語サポートと自動翻訳 (ベータ版): 海外からお申し込みですか?このシステムは、言語の不一致を厳密に防止してネイティブの多言語 CV をサポートし、データベースを動的にローカライズするための実験的な自動翻訳エンジンを備えています。
🏗️ 動的セクション ルーティング: 特定のアプリケーション用のオープンソース プロジェクトはありませんか? JSON 内の配列を省略するだけです。 Jinja2 エンジンはセクションを動的に非表示にし、厄介な空白を残さずに LaTeX ジオメトリを再計算します。
🐳 コンテナ化された再現性: 事前構成された VS Code DevC が付属しています

[切り捨てられた]

## Original Extract

Contribute to Gotili/EigenCV development by creating an account on GitHub.

GitHub - Gotili/EigenCV · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Gotili
/
EigenCV
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
44 Commits 44 Commits .devcontainer .devcontainer cv cv docs docs scripts scripts tests tests tools tools .cursorrules .cursorrules .gitignore .gitignore AI_START_HERE.md AI_START_HERE.md LICENSE LICENSE README.md README.md USER_GUIDE.md USER_GUIDE.md application_tracking.template.md application_tracking.template.md chatgpt_run.py chatgpt_run.py check_ats_score.py check_ats_score.py requirements.txt requirements.txt View all files Repository files navigation
🛡️ EigenCV: Zero-Trust Agentic Resume Pipeline
Stop letting ChatGPT hallucinate skills you don't have.
A production-grade Infrastructure-as-Code (IaC) pipeline for generating ATS-perfect, highly tailored LaTeX resumes without sacrificing your integrity.
🤯 Commercial AI Builders vs. EigenCV
The Industry Standard (Commercial AI Builders):
You tell an AI to "optimize my resume for this job." The AI treats your resume as a creative writing exercise. It quietly hallucinates skills, inflates job titles, and paraphrases your engineering achievements into generic HR buzzwords. The result is a PDF that beats the ATS but fails the technical interview because it's full of lies.
The EigenCV Approach (Zero-Trust):
We treat your career history as an immutable database. The AI is strictly an orchestration layer . It does not write your resume; it queries your database to pull the most relevant, pre-verified bullet points.
If the AI attempts to go rogue and hallucinate a missing skill into your profile to artificially boost your ATS score, the Python compiler's Lie Detector intercepts it and hard-crashes the build. Obvious lies never make it into the PDF.
🎨 Total Customization & Control
The sample above is just one configuration. EigenCV puts you in complete control:
Dynamic Sections: Not all sections are mandatory. Choose exactly what to include and reorder them instantly (e.g., move 'Education' to the top) by simply editing the comma-separated cvorder variable.
Change Layouts: Seamlessly swap between professional LaTeX templates (e.g., Awesome-CV , EigenCV-Modern ).
Brand Yourself: Inject custom corporate accent colors to match the company you're applying to.- Multilingual: Generate applications in English, German, or any other language using the built-in translation matrix.
🚀 How to Use EigenCV: Choose Your Path
Path 1: The "No-Code" Lifehack (For Non-Coders / ChatGPT Plus)
You don't need to know Python, LaTeX, or Git to use EigenCV. You can run the entire pipeline directly in ChatGPT.
Download this entire repository as a ZIP file.
Build your Master Database: Upload the ZIP to ChatGPT (requires Advanced Data Analysis) or Claude, along with ALL your old resumes. Tell the AI: "Extract all my career facts from these resumes and populate the EigenCV JSON database."
Apply to a Job: Paste the Job Description and tell the AI: "Please apply to this job using the instructions found in docs/AI_CLOUD_PROMPT.md ."
Download your PDF: ChatGPT will automatically match your database to the job, generate the JSON, and run the chatgpt_run.py wrapper script. Because we include a dedicated "Cloud-Safe" LaTeX template, ChatGPT will render the PDF directly in its sandbox and give you a download link!
(Fallback: If ChatGPT times out, it will still generate the .tex code. You can drag & drop that code into a free Overleaf account for instant rendering (or render the pdf locally using the pdflatex )).
Path 2: The Agentic Developer Route (CLI & IDE)
If you prefer terminal workflows, need maximum build speed, or require strict data privacy, run the pipeline locally.
Prerequisites: You need Python 3.11+ and a LaTeX distribution (e.g., TeX Live or MiKTeX). Alternatively, simply open this repo in VS Code and click "Reopen in Container" to use our pre-built Docker environment!
Install Dependencies: Run pip install -r requirements.txt .
Agent Setup:
For Maximum Speed: Open the repository using an Agentic CLI (like Antigravity / Claude Code ) or an Agentic IDE ( Cursor / Windsurf ) connected to their standard cloud APIs.
For 100 % Hardcore Privacy: Point your Agentic IDE to a local model (like Ollama or LM Studio ) so your data NEVER leaves your machine.
Build the DB: Tell the Agent: "Migrate my old CV. Follow AI_START_HERE.md ." to build your Zero-Trust database.
Apply: Paste a Job Description and say: "Apply to this job. Follow AI_START_HERE.md ."
Automation: The Agent will automatically route the prompts, generate the strict JSON, and execute the Python scripts locally to render your PDF and calculate your ATS score!
Most AI tools use a "generate and pray" approach. EigenCV uses Agentic Determinism . Here is how we guarantee a flawless, hallucination-free application:
The Immutable Master Database: You don't paste your resume into a chat window. Your career history lives offline as a structured JSON/Markdown database on your hard drive. Every achievement, project, and skill has a mathematically verifiable ID (e.g., proj_aws_migration ).
AI Orchestration (The Brain): When applying for a job, the AI reads the Job Description and acts as a strategic orchestrator. Instead of writing text, it executes a precise query against your database, returning only the IDs that maximize your ATS match score. It crafts a hyper-authentic Cover Letter based exclusively on your verified profile data.
Deterministic Compilation (The Muscle): The AI is entirely locked out of the rendering process. The EigenCV Python compiler takes the approved list of IDs, fetches the exact text from your offline database, and deterministically injects it into a stunning, ATS-optimized LaTeX template. Flawless layouts, no hallucinations.
sequenceDiagram
autonumber
actor U as 👤 User
participant LLM as 🧠 Agent
participant DB as 🗄️ Database
participant Py as 🐍 Python Tools
rect rgb(40, 40, 60)
Note over U,DB: Phase 0: Setup & Privacy
U->>Py: Run scrub_data.py
Py->>DB: Wipe private data
U->>LLM: Upload old CV
LLM->>DB: Build JSON Database
end
rect rgb(20, 40, 20)
Note over U,Py: Phase 1: Routing
U->>LLM: Provide Job Description
LLM->>DB: Query database
DB-->>LLM: Return valid IDs
LLM->>LLM: Generate build_config.json
end
rect rgb(60, 40, 20)
Note over LLM,Py: Phase 2: Compile
LLM->>Py: Execute cv_compiler.py
Note over Py,DB: 🏥 Healer auto-corrects IDs<br>🛡️ Lie Detector prevents hallucinations
Py->>DB: Fetch verified text
Py-->>LLM: Render PDF via Jinja2
end
rect rgb(40, 60, 40)
Note over LLM,Py: Phase 3: Verification
LLM->>Py: Execute check_ats_score.py
Py->>Py: Parse compiled PDF
Py-->>LLM: Return ATS Score
LLM-->>U: Present Final PDF & Score!
end
Loading
🚀 The "Lie Detector" in Action
+----------------------------- EigenCV Compiler ------------------------------+
| Compiling CV from build_config.json... |
+-----------------------------------------------------------------------------+
Using layout template: eigencv_resume.tex.j2, Locale: en
+--------------------------- Zero-Trust Violation ----------------------------+
| Zero-Trust Violation: You declared 'Rust' as a missing skill, but it was |
| hallucinated into the CV output! |
| You cannot artificially inject skills you do not have into free-text fields.|
+-----------------------------------------------------------------------------+
ValueError: Hallucinated skill detected: Rust
The user removes the hallucinated skill and recompiles:
+----------------------------- EigenCV Compiler ------------------------------+
| Compiling CV from build_config.json... |
+-----------------------------------------------------------------------------+
Successfully compiled CV to CV-Applicant-Google.tex
Auto-compiling PDFs with pdflatex...
Successfully compiled CV-Applicant-Google.pdf
ATS Match Score: 85.0 %
+-------------------------------------------------------------------+
| Category | Skills |
|-----------------+-------------------------------------------------|
| Missing (1) | Rust |
+-------------------------------------------------------------------+
[!] ATS Penalty Applied: 1 critical gap identified.
✨ Core USPs
🛡️ Zero-Trust & The EigenTruth Engine: Your career history lives in a static JSON database. If the LLM attempts to hallucinate a skill you don't have into your profile to artificially boost your ATS score, the compiler's EigenTruth Engine (our Lie Detector) catches it and hard-crashes the build.
🔒 Immutable Database: Your bullet points and skills are strictly IMMUTABLE . You can maintain them yourself or use LLMs to prep them, but within the EigenCV pipeline, the AI is only allowed to select them, never rewrite them.
✍️ Authentic Cover Letters: The AI uses your personal_dossier.md to write hyper-authentic cover letters based only on your real soft skills and hobbies, eliminating corporate fluff.
🎨 Corporate Auto-Coloring: The AI automatically deduces the target company's corporate identity and dynamically injects matching accent colors into the LaTeX output (or you can override it manually).
🧮 Advanced ATS Engine & Reality Checks: The post-compilation Python parser calculates a mathematically honest ATS keyword match score. Meanwhile, the AI Agent acts as a ruthless filter, estimating interview/offer probabilities and salary ranges based strictly on your verified skills, creating a realistic Probability Matrix.
🏥 Self-Healing IDs: If the LLM makes a minor typo when selecting an ID from your database (e.g., aws_mig instead of aws_migration ), the compiler's built-in rapidfuzz heuristics will auto-heal the ID, preventing brittle build crashes.
📄 Automated LaTeX Compilation: No more broken LaTeX parsing or missing brackets. The AI generates a strictly typed Pydantic JSON schema, deterministically compiled into beautiful Jinja2 LaTeX templates.
🌍 Multi-Language Support & Auto-Translation (Beta): Applying abroad? The system supports native multi-language CVs with strict language mismatch prevention, and features an experimental auto-translation engine to dynamically localize your database.
🏗️ Dynamic Section Routing: Don't have any open-source projects for a specific application? Simply omit the array in the JSON. The Jinja2 engine will dynamically hide the section and recalculate the LaTeX geometry without leaving awkward whitespace.
🐳 Containerized Reproducibility: Comes with a pre-configured VS Code DevC

[truncated]
