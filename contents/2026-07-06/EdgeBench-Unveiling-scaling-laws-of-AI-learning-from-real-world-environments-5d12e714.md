---
source: "https://github.com/ByteDance-Seed/EdgeBench"
hn_url: "https://news.ycombinator.com/item?id=48801308"
title: "EdgeBench: Unveiling scaling laws of (AI) learning from real-world environments"
article_title: "GitHub - ByteDance-Seed/EdgeBench: EdgeBench: Unveiling scaling laws of learning from real-world environments · GitHub"
author: "rguiscard"
captured_at: "2026-07-06T07:02:31Z"
capture_tool: "hn-digest"
hn_id: 48801308
score: 1
comments: 0
posted_at: "2026-07-06T06:31:39Z"
tags:
  - hacker-news
  - translated
---

# EdgeBench: Unveiling scaling laws of (AI) learning from real-world environments

- HN: [48801308](https://news.ycombinator.com/item?id=48801308)
- Source: [github.com](https://github.com/ByteDance-Seed/EdgeBench)
- Score: 1
- Comments: 0
- Posted: 2026-07-06T06:31:39Z

## Translation

タイトル: EdgeBench: 現実世界の環境から学習する (AI) のスケーリング則を明らかにする
記事のタイトル: GitHub - ByteDance-Seed/EdgeBench: EdgeBench: 実世界の環境から学習するスケーリングの法則を明らかにする · GitHub
説明: EdgeBench: 現実世界の環境から学習するスケーリングの法則を明らかにする - ByteDance-Seed/EdgeBench

記事本文:
GitHub - ByteDance-Seed/EdgeBench: EdgeBench: 現実世界の環境から学習するスケーリングの法則を明らかにする · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
バイトダンスシード
/
エッジベンチ
公共
通知
あなたはむ

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
9 Commits 9 Commits .github/ workflows .github/ workflows assets assets docs docs examples examples sforge sforge .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
EdgeBench は、自律 AI エージェントが現実世界の環境からどのように学習するかを評価するための 134 の実世界タスクのベンチマークです。 Instead of measuring one-shot performance, EdgeBench places agents in executable task environments with realistic, multi-level feedback and lets them iterate for 12+ hours per task — tracking the full trajectory of improvement, not just the final score. 51 のタスクと完全な評価フレームワークを公開します。
Analyzing ~38,000 hours of agent interaction on all 134 tasks, we find that performance follows a log-sigmoid scaling law as a function of interaction time ( $R^2 = 0.998$ ).詳細については、技術レポートを参照してください。
モデル
@2時間
@4時間
@6時間
@8h
@10h
@12時
クロード 作品4.8
39.0
45.7
48.1
49.8
50.9
51.3
GPT-5.5
36.8
42.1
44.5
46.3
47.6
48.4
GPT-5.4
29.7
34.0
36.5
38.0
38.9
39.3
GLM-5.1
26.0
30.4
32.9
34.9
36.5
37.4
DS-V4-プロ
23.3
27.1
29.0
29.9
30.9
31.0
カテゴリスコア @12h (134 タスク)
モデル
科学と機械学習
Systems & SE
最適化
知識
フォーマル
ゲーム
クロード 作品4.8
48.5
67.4
36.5
47.0
55.0
39.3
GPT-5.5
44.3
65.0
33.6
45.7
50.0
39.1
GPT-5.4
33.5
54.1
27.9
38.8
40.8
29.0
GLM-5.1
33.8
50.9
26.4
43.5
24.6
29.3
DS-V4-Pro
30.0
43.0
21.5
37.0
14.1
16.9
オープンソースのサブセット (51 タスク)
モデル
@2時間
@4h
@6時間
@8h
@10h
@12h
クロード 作品4.8
33.2
38.5
40.8
42.1
43.3
44.2
GPT-5.5
31.2
36.0
38.2
40.3
42.1
43

.1
GPT-5.4
25.0
28.2
30.3
32.1
33.3
34.2
GLM-5.1
21.4
24.2
26.8
28.2
29.1
30.4
DS-V4-プロ
17.1
21.1
22.9
23.8
25.1
25.7
Category Scores @12h (51 tasks)
モデル
科学と機械学習
Systems & SE
最適化
知識
フォーマル
ゲーム
クロード 作品4.8
38.9
62.0
38.2
38.7
40.9
39.3
GPT-5.5
33.2
60.5
32.3
38.4
49.0
39.1
GPT-5.4
24.6
50.1
29.9
31.6
30.2
29.0
GLM-5.1
26.8
43.6
26.7
31.0
19.9
29.3
DS-V4-プロ
31.1
37.6
24.1
33.2
12.7
16.9
Per-Task Scores by Time Budget (51 tasks)
Each model cell reports scores at @2h / @4h / @6h / @8h / @10h / @12h . Missing valid results are shown as — .
EdgeBench contains 134 realistic, diverse tasks spanning six capability categories, of which 51 are publicly released . Each task is designed as a day-scale challenge with a performance ceiling high enough that no current agent can saturate it. Recorded human expert effort averages 57.2 hours per task (up to 320 hours).
EdgeBench is powered by SForge , a two-container evaluation harness built for long-horizon agent evaluation.各タスクは分離された作業として具体化され、Docker イメージを判断します。エージェントは作業環境のみを参照し、非表示のテストは一時的な判断コンテナで実行されます。
Two-container isolation — work and judge environments are fully separated, preventing evaluation hacking at its root
Iterative evaluation with feedback — agents don't submit once at the end for a one-shot score;代わりに、実行全体を通して提出し、詳細なフィードバック (合格率、テストの不合格、スコア) を受け取り、タイムアウトまで閉ループで改善します。すべての提出で最良の結果が最終スコアとなります。
長期的な実行 - 停止フックによりエージェントの早期終了が防止され、自動再開により一時的な障害から回復され、Kubernetes バックエンドにより大規模な並列実行が可能になります。
# Install (requires Docker Engin

e running on a Linux host)
pip インストール sforge
# 1. Download task definitions
sforge fetch-tasks edgebench
# 2. Pull pre-built Docker images
sforge pull --task ad_placement_optimization --registry seededge
# 3. Start judge server (separate terminal)
サーブを鍛錬する
# 4. エージェントを実行する
SFORGE_AGENT_API_KEY= " sk-xxx " \
sforge run --task ad_placement_optimization --agent claude-code \
--モデル " クロード-opus-4-8[1m] " --timeout 43200 --run-idedgebench-001
段階的な例:
ローカル Docker 上の単一タスク — Docker を使用して 1 つのタスクをエンドツーエンドで実行します
Kubernetes 上のすべてのタスク — 実験 YAML を使用して K8s クラスター上で完全なスイートを実行します
Evaluating your own model / agent:
独自のモデル — 組み込みの Claude Code および Codex スキャフォールドは、互換性のある API エンドポイントで動作します。エンドポイントで SFORGE_AGENT_API_BASE_URL を指定し、 SFORGE_AGENT_API_KEY を介してキーを設定し、 --model を介してモデル名を渡します。 「サポートされているエージェント」を参照してください。
独自のエージェント スキャフォールド — sforge/harness/agent/ (インストールおよび起動方法を宣言する小さなエージェント サブクラス) の下に新しいエージェントを追加し、それをファクトリに登録し、 --agent <your-agent> で実行します。 「カスタム エージェント」を参照してください。
完全なドキュメント: bytedance-seed.github.io/EdgeBench
EdgeBench が研究に役立つと思われる場合は、当社の技術レポートを引用してください。
@その他 {edgebench2026 ,
title = { EdgeBench: 現実世界の環境から学習するスケーリングの法則を明らかにする } ,
著者 = { Deyao Zhu と Xin Zhou と Shengling Qin と Xuekai Zhu と Hangliang Ding と Shu Zhong とその他 } ,
年 = { 2026 } 、
url = { https://edge-bench.org/paper.pdf } ,
}
ライセンス
EdgeBench タスク (タスク データセット) は CC BY 4.0 に基づいてリリースされます。
SForge (評価ハーネス コード) は、Apache License 2.0 に基づいてリリースされています。
134 個のタスク スイート全体を評価するには、zhongshu@bytedance.com までご連絡ください。
E

dgeBench: 現実世界の環境から学習するスケーリングの法則を明らかにする
Apache-2.0ライセンス
貢献する
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
5
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

EdgeBench: Unveiling scaling laws of learning from real-world environments - ByteDance-Seed/EdgeBench

GitHub - ByteDance-Seed/EdgeBench: EdgeBench: Unveiling scaling laws of learning from real-world environments · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
ByteDance-Seed
/
EdgeBench
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
9 Commits 9 Commits .github/ workflows .github/ workflows assets assets docs docs examples examples sforge sforge .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml View all files Repository files navigation
EdgeBench is a benchmark of 134 real-world tasks for evaluating how autonomous AI agents learn from real-world environments . Instead of measuring one-shot performance, EdgeBench places agents in executable task environments with realistic, multi-level feedback and lets them iterate for 12+ hours per task — tracking the full trajectory of improvement, not just the final score. We publicly release 51 tasks along with the full evaluation framework.
Analyzing ~38,000 hours of agent interaction on all 134 tasks, we find that performance follows a log-sigmoid scaling law as a function of interaction time ( $R^2 = 0.998$ ). See the tech report for details.
Model
@2h
@4h
@6h
@8h
@10h
@12h
Claude Opus 4.8
39.0
45.7
48.1
49.8
50.9
51.3
GPT-5.5
36.8
42.1
44.5
46.3
47.6
48.4
GPT-5.4
29.7
34.0
36.5
38.0
38.9
39.3
GLM-5.1
26.0
30.4
32.9
34.9
36.5
37.4
DS-V4-Pro
23.3
27.1
29.0
29.9
30.9
31.0
Category Scores @12h (134 tasks)
Model
Scientific & ML
Systems & SE
Optimization
Knowledge
Formal
Games
Claude Opus 4.8
48.5
67.4
36.5
47.0
55.0
39.3
GPT-5.5
44.3
65.0
33.6
45.7
50.0
39.1
GPT-5.4
33.5
54.1
27.9
38.8
40.8
29.0
GLM-5.1
33.8
50.9
26.4
43.5
24.6
29.3
DS-V4-Pro
30.0
43.0
21.5
37.0
14.1
16.9
Open-Source Subset (51 tasks)
Model
@2h
@4h
@6h
@8h
@10h
@12h
Claude Opus 4.8
33.2
38.5
40.8
42.1
43.3
44.2
GPT-5.5
31.2
36.0
38.2
40.3
42.1
43.1
GPT-5.4
25.0
28.2
30.3
32.1
33.3
34.2
GLM-5.1
21.4
24.2
26.8
28.2
29.1
30.4
DS-V4-Pro
17.1
21.1
22.9
23.8
25.1
25.7
Category Scores @12h (51 tasks)
Model
Scientific & ML
Systems & SE
Optimization
Knowledge
Formal
Games
Claude Opus 4.8
38.9
62.0
38.2
38.7
40.9
39.3
GPT-5.5
33.2
60.5
32.3
38.4
49.0
39.1
GPT-5.4
24.6
50.1
29.9
31.6
30.2
29.0
GLM-5.1
26.8
43.6
26.7
31.0
19.9
29.3
DS-V4-Pro
31.1
37.6
24.1
33.2
12.7
16.9
Per-Task Scores by Time Budget (51 tasks)
Each model cell reports scores at @2h / @4h / @6h / @8h / @10h / @12h . Missing valid results are shown as — .
EdgeBench contains 134 realistic, diverse tasks spanning six capability categories, of which 51 are publicly released . Each task is designed as a day-scale challenge with a performance ceiling high enough that no current agent can saturate it. Recorded human expert effort averages 57.2 hours per task (up to 320 hours).
EdgeBench is powered by SForge , a two-container evaluation harness built for long-horizon agent evaluation. Each task materializes as isolated work and judge Docker images — the agent only sees the work environment, while hidden tests run in ephemeral judge containers.
Two-container isolation — work and judge environments are fully separated, preventing evaluation hacking at its root
Iterative evaluation with feedback — agents don't submit once at the end for a one-shot score; instead they submit throughout the run, receive granular feedback (pass rates, failing tests, scores), and improve in a closed loop until timeout — the best result across all submissions is the final score
Long-horizon execution — stop hooks prevent premature agent exit, auto-resume recovers from transient failures, and the Kubernetes backend enables parallel runs at scale
# Install (requires Docker Engine running on a Linux host)
pip install sforge
# 1. Download task definitions
sforge fetch-tasks edgebench
# 2. Pull pre-built Docker images
sforge pull --task ad_placement_optimization --registry seededge
# 3. Start judge server (separate terminal)
sforge serve
# 4. Run an agent
SFORGE_AGENT_API_KEY= " sk-xxx " \
sforge run --task ad_placement_optimization --agent claude-code \
--model " claude-opus-4-8[1m] " --timeout 43200 --run-id edgebench-001
Step-by-step examples:
Single task on local Docker — run one task end-to-end with Docker
All tasks on Kubernetes — run the full suite on a K8s cluster with experiment YAML
Evaluating your own model / agent:
Your own model — the built-in Claude Code and Codex scaffolds work with any compatible API endpoint: point SFORGE_AGENT_API_BASE_URL at your endpoint, set your key via SFORGE_AGENT_API_KEY , and pass your model name via --model . See Supported Agents .
Your own agent scaffold — just add a new agent under sforge/harness/agent/ (a small Agent subclass declaring how to install and launch it) and register it in the factory, then run with --agent <your-agent> . See Custom Agents .
Full documentation: bytedance-seed.github.io/EdgeBench
If you find EdgeBench useful in your research, please cite our tech report:
@misc { edgebench2026 ,
title = { EdgeBench: Unveiling Scaling Laws of Learning from Real-World Environments } ,
author = { Deyao Zhu and Xin Zhou and Shengling Qin and Xuekai Zhu and Hangliang Ding and Shu Zhong and others } ,
year = { 2026 } ,
url = { https://edge-bench.org/paper.pdf } ,
}
License
EdgeBench Tasks (task datasets) are released under CC BY 4.0 .
SForge (evaluation harness code) is released under the Apache License 2.0 .
To evaluate on the full 134-task suite, please contact zhongshu@bytedance.com .
EdgeBench: Unveiling scaling laws of learning from real-world environments
Apache-2.0 license
Contributing
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
5
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
