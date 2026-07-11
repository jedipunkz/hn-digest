---
source: "https://blog.grandimam.com/posts/how-hacker-rank-scores-engineers/"
hn_url: "https://news.ycombinator.com/item?id=48873620"
title: "HackerRank's LLM-Based Hiring Agent"
article_title: "How HackerRank Scores Engineers | Grandimam"
author: "grandimam"
captured_at: "2026-07-11T17:46:37Z"
capture_tool: "hn-digest"
hn_id: 48873620
score: 1
comments: 0
posted_at: "2026-07-11T16:55:18Z"
tags:
  - hacker-news
  - translated
---

# HackerRank's LLM-Based Hiring Agent

- HN: [48873620](https://news.ycombinator.com/item?id=48873620)
- Source: [blog.grandimam.com](https://blog.grandimam.com/posts/how-hacker-rank-scores-engineers/)
- Score: 1
- Comments: 0
- Posted: 2026-07-11T16:55:18Z

## Translation

タイトル: HackerRank の LLM ベースの採用エージェント
記事のタイトル: HackerRank がエンジニアを採点する方法 |グランディマム
説明: LLM スコアリング ツールの構築で最も難しい部分は、

記事本文:
← グランディマム
HackerRank がエンジニアを採点する方法
私は LLM を通じて目標設定、KPI、フィードバック評価を行ってきました。最も困難だったのは、スコアリングのメカニズムを設計することでした。 HackerRank が Hiring Agent をオープンソース化したとき、私は彼らが同じ問題をどのように解決したのか、そして何がルーブリック エンコードにバイアスを与えているのかを理解したいと思いました。
Hiring Agent は、PDF を解析し、GitHub とブログ データで強化してスコアを返す LLM ベースの履歴書スコアリング ツールです。実行は簡単です。リポジトリのクローンを作成し、PDF を指定します。
$ git clone https://github.com/interviewstreet/hiring-agent
$ cd 採用エージェント
$ python3 スコア.py /examples/java-engineer.pdf
このツールは構造化されたスコアを返します。
=========================================
📊 評価の履歴書: シニア Java エンジニア
=========================================
🎯 総合スコア: 71.0/100
🌐 オープンソース: 12/35
🚀 セルフプロジェクト: 18/30
🏢 制作経験: 25/25
💻 技術スキル: 8/10
✅ 主な強み:
------------------------------------------
1. マイクロサービスアーキテクチャ
2. イベント駆動型システム
3. クラウドテクノロジー
4. データベースの最適化
5. API設計
🔧 改善の余地がある領域:
------------------------------------------
1. オープンソースへの貢献の欠如
2. 限定的な GitHub アクティビティ
3. 公開されているプロジェクトはありません
私は最初にシニア Java エンジニアの履歴書を提出しました。プロフィールには制作経験と技術スキルに関してほぼ完璧なスコアが付いていますが、ツールによる候補者の評価は依然として 71 点でした。私は、このツールがどのようにスコアリングを行っているのかをより深く理解したいと思いました。
これは、PDF が PDFHandler に渡される core.py で始まります。
pdf_handler = PDFHandler ()
再開データ = pdf_handler 。 extract_json_from_pdf ( pdf_path )
extract_json_from_pdf 内で PDF が解析されます

d PyMuPDF を使用し、Markdown に変換しました。
履歴書テキスト = to_markdown (
ドクター、
ページ = ページ 、
)
その後、Markdown コンテンツは構造化された JSON、つまり履歴書の内容を表す内部スキーマ (基本、仕事、教育、スキル、プロジェクト、賞) に正規化されます。正規化ステップは、ツールがレイアウトに関係なく履歴書を取り込み、厳密な内部形式で表現できるようにするため、重要です。現在、履歴書を PDF から Markdown に解析することは確実な手順ではありません。複数列セクションの場合、Markdown への直接変換は信頼できません。
この履歴書の例を見てみましょう。
Markdown パーサーは左から右に読み取るため、コンテンツは次のように誤って抽出されます。
コーディング、フレームワーク、メッセージング、データベース、可観測性、コンテナ、パターン、クラウド、Java、Python…
LLM がこの状態からどの程度回復するかは、使用するモデルによって異なります。より具体的な解決策は、履歴書をスクリーンショットして、最初に Markdown に変換するのではなく、モデルに直接解析させることです。これは、Claude/Codex が Web サイトを解析する方法と同様です。次に、特殊なプロンプトを使用してセクションが個別に抽出されます。各プロンプトはモデルに対する LLM 呼び出しです。
ほとんどすべてのインテリジェンスはプロンプトであり、prompts/templates/ ディレクトリで入手できます。 Python はオーケストレーションと配管にのみ使用されます。
"基本": "基本.jinja",
"work": "work.jinja",
"教育": "education.jinja",
"スキル": "スキル.神社",
"プロジェクト": "projects.jinja",
"賞": "賞.jinja",
"system_message": "system_message.jinja",
"github_project_selection": "github_project_selection.jinja",
"resume_evaluation_criteria": "resume_evaluation_criteria.jinja",
"resume_evaluation_system_message": "resume_evaluation_system_message.jinja",
resume_evaluation_system_message.jinja とresume_evaluation_criteria はスコアの評価に役立ちます。オット

彼女のプロンプトは主に履歴書の解析機能を目的としています。 resume_evaluation_criteria の採点基準は、オープンソースのプロジェクトや貢献を行っている個人に大きく偏っています。
### オープンソース (0-35 ポイント)
**高スコア (25 ～ 35 ポイント):**
- 人気のオープンソース プロジェクトへの貢献 (1000 個以上のスター)
- 有名なプロジェクトへの多大な貢献
- Google Summer of Code (GSoC) への参加
- コミュニティへの実質的な関与
どのプロジェクトがオープンソースとみなされ、どのプロジェクトがそうでないかについて言及する明確なルールがあります。
**重要なルール:**
- 個人の GitHub リポジトリを持つことは、オープンソースへの貢献にはなりません
- 真のオープンソースへの貢献とは、他の人のプロジェクトに貢献することを意味します
- GitHub データがすべてのプロジェクトが「self_project」タイプであることを示している場合、オープンソース スコアは 10 ポイント以下でなければなりません
そして私が最も興味を持ったのは次のルールでした。
**スタートアップの経験に関する特別な考慮事項**: スタートアップにおける創業者の役割、共同創設者の役割、または初期段階のエンジニアの役割 (最初の 10 ～ 20 人の従業員) には、並外れたイニシアチブ、技術的なリーダーシップ、および製品をゼロから構築する能力を示すため、追加ポイントを与えます。
バイアスをテストする
プロンプトには、創設者と初期段階の役割に報酬を与えるようにと記載されています。私が知りたかったのは、実際にそうなのかということです。同じ履歴書を使用し、役職名だけを変更しました。 3 つのバージョン:
創業エンジニア - 同じ履歴書、肩書を交換
共同創設者 / CTO - 同じ履歴書、肩書きが再び交換
それ以外はすべて同じままでした。同じスキル、同じプロジェクト、同じ経験です。唯一の変数は役職でした。
結果は圧倒的なものでした。いくつかのイテレーションでは、プロダクション スコアが 25 点中 22 から 23 に変化しました。 CTOの称号も同じ範囲に収まった。異なるタイトルで複数回実行しましたが、プロダクション スコアは一貫していました

私が何を置いたかに関係なく、20〜23の間に座っていました。
これは 2 つのことを物語っています。まず、LLM はタイトルの変更に耐性があります。第二に、ルーブリックでスタートアップの役割が強調されているのは、強制というよりも願望です。プロンプトでは追加のポイントを与えるように指示されていますが、モデルはそれらを一貫して適用しません。プロンプト ルールは提案であり、保証ではありません。
基準を詳しく調べたところ、次のような結果も得られました。
- スタートアップの創設者/共同創設者の経験に対して +3 ～ 5 ポイント
- 初期段階のエンジニア経験に対して +2 ～ 3 ポイント (スタートアップの最初の 10 ～ 20 人の従業員)
- ポートフォリオ Web サイトの +2 ポイント (Basics.url 内の GitHub URL)
- LinkedIn プロフィールで +1 ポイント
LinkedIn ユーザーであることは 1 ポイントの価値があります。 GSoC 参加者であることは 5 の価値があります。プロンプトとは別に、リポジトリの人気をチェックするルールが明示的にコード化されています。
repos_data のリポジトリの場合:
もしレポがあれば。 get (「fork」) とリポジトリ。 get ( " forks_count " , 0 ) < 5 :
続ける
フォークが 5 つ未満のフォークを削除し、人気のあるリポジトリのドライブバイ フォークを除外します。これらはルールとして述べられた意見です。
Hiring Agent は、適切に設計されたプロトタイプです。問題は、オープンソースがエンジニアの価値の 35% であると誰かが決めたことです。 4 つのコミットが「意味のある」貢献のしきい値です。 LinkedIn プロフィールはちょうど 1 ポイントの価値があると誰かが判断しました。これらはルールとして定められた判断です。
これは私が独自のツールで解決しようとしていた問題と同じです。 LLM ベースのスコアリングの難しい部分は、モデル、解析、パイプラインではありません。何が重要かを決めるのです。プロンプトで定義すると、値の判断が行われたことになり、モデルはそれを強制しようとします。
gitの歴史も興味深いです。活発な開発は 77 日間 (2025 年 7 月 29 日から 10 月 15 日まで) に及び、その後 231 日間の沈黙期間が続き、その後 2026 年 6 月にいくつかのクリーンアップ コミットが行われてからオープンソース化されます。

してる。コミット パターンは、継続的な取り組みではなく、社内ハッカソンやインターンの作業が一気に行われたように見えます。
プロジェクトをオープンソース化している企業はこの分野のプレーヤーの 1 つであるため、これは重要です。定義された基準: オープンソースが 35%、技術スキルが 10% は誰かの初稿のように感じられ、徹底的に検証された採用フレームワークではありません。プロジェクトには矛盾点もあり、これらはプロトタイプの大まかな部分であり、実稼働システムではないことが強調されています。このツールは、何千人もの実際の候補者に対してテストされ、採用チームによって調整される前に、LLM ベースの採用を試みる初めての試みであるため、慎重に使用する必要があります。

## Original Extract

The hardest part of building an LLM scoring tool isn

← Grandimam
How HackerRank Scores Engineers
I have been doing my goal-setting, KPIs, and feedback assessments through LLMs. The hardest part was designing the scoring mechanism. When HackerRank open-sourced their Hiring Agent , I wanted to understand how they had solved the same problem — and what biases their rubric encodes.
Hiring Agent is an LLM-based resume scoring tool that parses a PDF, enriches it with GitHub and blog data, and returns a score. It’s straightforward to run — clone the repo, point it at a PDF:
$ git clone https://github.com/interviewstreet/hiring-agent
$ cd hiring-agent
$ python3 score.py /examples/java-engineer.pdf
The tool returns a structured score:
==========================================
📊 RESUME EVALUATION: Senior Java Engineer
==========================================
🎯 OVERALL SCORE: 71.0/100
🌐 Open Source: 12/35
🚀 Self Projects: 18/30
🏢 Production Experience: 25/25
💻 Technical Skills: 8/10
✅ KEY STRENGTHS:
------------------------------------------
1. Microservices Architecture
2. Event-Driven Systems
3. Cloud Technologies
4. Database Optimization
5. API Design
🔧 AREAS FOR IMPROVEMENT:
------------------------------------------
1. Lack of Open Source Contributions
2. Limited GitHub Activity
3. No Publicly Available Projects
I first provided it with a Senior Java Engineer resume. Though the profile has a near-perfect score on production experience and technical skills, the tool still evaluated the candidate at 71. I wanted to get a deeper understanding of how this tool was doing the scoring:
It starts in score.py where the PDF is handed off to a PDFHandler :
pdf_handler = PDFHandler ()
resume_data = pdf_handler . extract_json_from_pdf ( pdf_path )
Inside extract_json_from_pdf, the PDF is parsed using PyMuPDF and converted to Markdown:
resume_text = to_markdown (
doc ,
pages = pages ,
)
The Markdown content is then normalised into structured JSON - an internal schema (basics, work, education, skills, projects, awards) that represents the content of the resume. The normalisation step is important as it allows the tool to take in any resume regardless of the layout and represent it into a strict internal format. Now, parsing the resume from PDF to Markdown is not a foolproof step, in case of multi-column sections direct Markdown conversion is unreliable.
Let’s take an example of this resume:
The Markdown parser reads left-to-right, so the content is incorrectly extracted as:
Coding, Framework, Messaging, Database, Observability, Containers, Patterns, Cloud, Java, Python…
How well the LLM recovers from this depends on the model used. A more concrete solution would be to screenshot the resume and let the model parse it directly, rather than converting to Markdown first - similar to how Claude/Codex parses websites. The sections are then extracted individually using specialized prompts, each being an LLM call to the model.
Almost all the intelligence is prompts, available in prompts/templates/ directory. Python is merely used in orchestration and plumbing.
"basics": "basics.jinja",
"work": "work.jinja",
"education": "education.jinja",
"skills": "skills.jinja",
"projects": "projects.jinja",
"awards": "awards.jinja",
"system_message": "system_message.jinja",
"github_project_selection": "github_project_selection.jinja",
"resume_evaluation_criteria": "resume_evaluation_criteria.jinja",
"resume_evaluation_system_message": "resume_evaluation_system_message.jinja",
resume_evaluation_system_message.jinja and resume_evaluation_criteria help evaluate the scoring. The other prompts are primarily for resume parsing capability. The scoring criteria in resume_evaluation_criteria is heavily biased towards individuals with open-source projects and contributions:
### Open Source (0-35 points)
**HIGH SCORES (25-35 points):**
- Contributions to popular open source projects (1000+ stars)
- Significant contributions to well-known projects
- Google Summer of Code (GSoC) participation
- Substantial community involvement
There are explicit rules that mention what projects are considered open-source and what are not.
**CRITICAL RULES:**
- Having personal GitHub repositories does NOT constitute open source contribution
- True open source contribution means contributing to OTHER people's projects
- When GitHub data shows all projects are 'self_project' type, open source score MUST be 10 points or less
And one that interested me most was the following rule:
**SPECIAL CONSIDERATION FOR STARTUP EXPERIENCE**: Give extra points for founder roles, co-founder positions, or early-stage engineer roles (first 10-20 employees) at startups, as these demonstrate exceptional initiative, technical leadership, and ability to build products from scratch.
Testing the Bias
The prompt says to reward founder and early-stage roles. I wanted to know: does it actually? I took the same resume and changed only the job titles. Three versions:
Founding Engineer - same resume, title swapped
Co-founder / CTO - same resume, title swapped again
Everything else stayed identical: same skills, same projects, same experience. The only variable was the job title.
The results were underwhelming. The production score moved from 22 to 23 out of 25 in some iterations. The CTO title landed in the same range. I ran it multiple times with different titles, and the production score consistently sat between 20-23 regardless of what I put.
This tells two things. First, the LLM is resistant to title changes. Second, the rubric’s emphasis on startup roles is more aspiration than enforcement. The prompt says to give extra points, but the model doesn’t consistently apply them. A prompt rule is a suggestion, not a guarantee.
A deeper look into the criteria also rewarded the following:
- +3-5 points for startup founder/co-founder experience
- +2-3 points for early-stage engineer experience (first 10-20 employees at a startup)
- +2 points for portfolio website (GitHub URL in basics.url)
- +1 point for LinkedIn profile
Being a LinkedIn user is worth 1 point. Being a GSoC participant is worth 5. Now, apart from prompts there are rules explicitly coded wherein it checks for repo popularity:
for repo in repos_data :
if repo . get ( " fork " ) and repo . get ( " forks_count " , 0 ) < 5 :
continue
Drops forks with fewer than 5 forks, filtering out drive-by forks of popular repos. These are opinions stated as rules.
The Hiring Agent is a decently-engineered prototype. The problem is someone decided open source is 35% of an engineer’s value; four commits is the threshold for “meaningful” contribution. Someone decided a LinkedIn profile is worth exactly 1 point. These are judgment stated as rules.
This is the same problem I was trying to solve in my own tool. The hard part of LLM-based scoring isn’t the model, the parsing, or the pipeline. It’s deciding what matters. Once you define it in the prompt, you have made a value judgment and the model will try to enforce it.
The git history is also interesting. The active development spans 77 days (July 29 to October 15, 2025) followed by a 231-day silence, then a handful of cleanup commits in June 2026 before open-sourcing. The commit patterns look like a burst of internal hackathon or intern work, not a sustained effort.
This matters because the company that is open-sourcing the project is one of the players in the domain. The criteria defined: open source at 35%, technical skills at 10% feels like someone’s first draft and not a thoroughly validated hiring framework. There are inconsistencies also in the project that highlight that these are the rough edges of a prototype, not a production system. This tool must be used with caution as it’s a first attempt at LLM-based hiring looks like, before it’s been tested against thousands of real candidates and calibrated by recruiting teams.
