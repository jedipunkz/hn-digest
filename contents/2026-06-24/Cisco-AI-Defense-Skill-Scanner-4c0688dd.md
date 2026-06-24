---
source: "https://github.com/cisco-ai-defense/skill-scanner"
hn_url: "https://news.ycombinator.com/item?id=48656076"
title: "Cisco AI Defense Skill Scanner"
article_title: "GitHub - cisco-ai-defense/skill-scanner: Security Scanner for Agent Skills · GitHub"
author: "chha"
captured_at: "2026-06-24T07:07:46Z"
capture_tool: "hn-digest"
hn_id: 48656076
score: 2
comments: 0
posted_at: "2026-06-24T06:44:16Z"
tags:
  - hacker-news
  - translated
---

# Cisco AI Defense Skill Scanner

- HN: [48656076](https://news.ycombinator.com/item?id=48656076)
- Source: [github.com](https://github.com/cisco-ai-defense/skill-scanner)
- Score: 2
- Comments: 0
- Posted: 2026-06-24T06:44:16Z

## Translation

タイトル: Cisco AI 防御スキル スキャナー
記事のタイトル: GitHub - cisco-ai-defense/skill-scanner: エージェント スキル用のセキュリティ スキャナー · GitHub
説明: エージェント スキル用のセキュリティ スキャナー。 GitHub でアカウントを作成して、cisco-ai-defense/skill-scanner の開発に貢献してください。

記事本文:
GitHub - cisco-ai-defense/skill-scanner: エージェント スキル用のセキュリティ スキャナー · GitHub
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
シスコ-ai-ディフェンス
/
スキルスキャナー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
シスコ-アイ-ディフェ

NSE/スキルスキャナー
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
61 コミット 61 コミット .cursor/ rules .cursor/ rules .devin .devin .github .github .windsurf/ rules .windsurf/ rules Formula Formula docs docs evals evals 例 例 スクリプト スクリプト skill_scanner skill_scanner テスト テスト .env.example .env.example .gitignore .gitignore .gitleaksignore .gitleaksignore .pre-commit-config.yaml .pre-commit-config.yaml .pre-commit-hooks.yaml .pre-commit-hooks.yaml CODEOWNERS CODEOWNERS CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md FEATURE.md FEATURE.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md TESTING.md TESTING.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェント スキル向けのベストエフォート型セキュリティ スキャナーで、プロンプト インジェクション、データ引き出し、悪意のあるコード パターンを検出します。パターンベースの検出 (YAML + YARA)、LLM-as-a-judge、および動作データフロー分析を組み合わせて、誤検知を最小限に抑えながら、潜在的な脅威の検出範囲を最大化します。
重要: このスキャナーはベストエフォート型の検出を提供するものであり、包括的または完全な範囲をカバーするものではありません。スキャンで結果が返されなかった場合でも、スキルにすべての脅威がないことが保証されるわけではありません。以下の範囲と制限を参照してください。
エージェント スキル仕様に従って、OpenAI Codex スキルとカーソル エージェント スキル形式をサポートします。 --lenient を使用すると、Claude Code .claude/commands/*.md やフラット マークダウン スキル リポジトリなどの非標準形式もスキャンします。
マルチエンジン検出 - 静的分析、動作データフロー、LLM セマンティック分析、および階層化されたベストエフォート型のカバレッジを実現するクラウドベースのスキャン
誤検知フィルタリング - メタアナライザーは検出能力を維持しながらノイズを大幅に低減します

能力
CI/CD Ready - GitHub コード スキャン用の SARIF 出力、再利用可能な GitHub Actions ワークフロー、ビルド失敗時の終了コード
プレコミットフック - 各コミット前にスキルをスキャンするための標準のプレコミットフレームワーク統合
拡張可能 - カスタム アナライザー用のプラグイン アーキテクチャ
Cisco AI Discord に参加して、議論したり、フィードバックを共有したり、チームとつながりましょう。
Skill Scanner は検出ツールです。既知のリスク パターンと可能性の高いリスク パターンを特定しますが、セキュリティを保証するものではありません。
所見なし≠ リスクなし。スキャンで「検出結果なし」が返された場合は、既知の脅威パターンが検出されなかったことを示します。スキルが安全であること、無害であること、または脆弱性がないことを保証するものではありません。
補償内容は本質的に不完全です。このスキャナーは、シグネチャ ベースの検出、LLM ベースのセマンティック分析、動作データフロー分析、オプションのクラウド サービス、構成可能なルール パックを組み合わせています。このアプローチによりカバレッジは向上しますが、あらゆる手法、特に新規攻撃やゼロデイ攻撃を検出できる自動ツールはありません。
偽陽性と偽陰性が発生する可能性があります。コンセンサス モードとメタ分析によりノイズは軽減されますが、誤った分類をすべて排除できる構成はありません。リスク許容度に合わせてスキャン ポリシーを調整します。
人間によるレビューは引き続き不可欠です。自動スキャンは、多層防御戦略のコンポーネントの 1 つです。高リスクまたは運用環境の展開では、スキャナーの結果を手動のコード レビューや脅威モデリングと組み合わせる必要があります。
ガイド
説明
クイックスタート
5分以内に始めましょう
建築
システム設計とコンポーネント
脅威分類法
例を含む完全な AITech 脅威分類法
LLM アナライザー
LLM の構成と使用法
メタアナライザー
誤検知フィルタリングと優先順位付け
行動分析装置
データフロー分析の詳細
スキャンポリシー
カスタム ポリシー、プリセット、チューニング ガイド
ポリシーのクイックリファレンス
コンパクトリファレンス

ポリシーセクションとノブ用
ルールのオーサリング
署名、YARA、Python ルールを追加する方法
GitHub アクション
CI/CD統合のための再利用可能なワークフロー
APIリファレンス
REST API ドキュメント
開発ガイド
貢献と開発のセットアップ
インストール
前提条件: Python 3.10+ および uv (推奨) または pip
# UV の使用 (推奨)
uv pip インストール cisco-ai-skill-scanner
# pip を使用する
pip インストール cisco-ai-skill-scanner
クラウドプロバイダーの追加機能
# AWS Bedrock のサポート
pip インストール cisco-ai-skill-scanner[bedrock]
# Google AI Studio / Gemini のサポート
pip インストール cisco-ai-skill-scanner[google]
# Google Vertex AI のサポート
pip インストール cisco-ai-skill-scanner[頂点]
# Azure OpenAI サポート
pip インストール cisco-ai-skill-scanner[azure]
# すべてのクラウドプロバイダー
pip インストール cisco-ai-skill-scanner[すべて]
クイックスタート
# LLM アナライザーとメタアナライザーの場合
import SKILL_SCANNER_LLM_API_KEY= " your_api_key "
エクスポート SKILL_SCANNER_LLM_MODEL= " クロード-3-5-sonnet-20241022 "
# VirusTotal バイナリ スキャンの場合
import VIRUSTOTAL_API_KEY= " your_virustotal_api_key "
# Cisco AI 防御の場合
import AI_DEFENSE_API_KEY= " your_aidefense_api_key "
インタラクティブウィザード
どのフラグを使用すればよいかわからない場合は、引数を指定せずに skill-scanner を実行して、対話型ウィザードを起動します。
スキルスキャナー
ウィザードでは、スキャン ターゲット、アナライザー、ポリシー、出力形式を選択する手順が示され、実行する前に組み立てられたコマンドが表示されます。 CLI の学習に最適です。
# 単一のスキルをスキャンします (コア アナライザー: 静的 + バイトコード + パイプライン)
スキルスキャナスキャン /path/to/skill
# 動作アナライザーによるスキャン (データフロー分析)
skill-scanner scan /path/to/skill --use-behavioral
# すべてのエンジンでスキャン
skill-scanner scan /path/to/skill --use-behavioral --use-llm --use-aidefense
# メタアナライザーでスキャンして誤検知フィルタリングを行う
スキルスキャナスキャン /path/to/

スキル --use-llm --enable-meta
# トリガーアナライザーでスキャンして曖昧な説明をチェック
スキルスキャナスキャン /path/to/skill --use-trigger
# LLM アナライザーを複数回実行し、多数決で合意された結果を維持する
skill-scanner scan /path/to/skill --use-llm --llm-consensus-runs 3
# 複数のスキルを再帰的にスキャンします
skill-scanner scan-all /path/to/skills --recursive --use-behavioral
# クロススキルの重複検出で複数のスキルをスキャン
skill-scanner scan-all /path/to/skills --recursive --check-overlap
# 寛容モード: 失敗する代わりに不正なスキルを許容します
skill-scanner scan /path/to/skill --lenient
skill-scanner scan-all /path/to/skills --recursive --lenient
# 非標準スキル形式の寛容モード (SKILL.md は必要ありません)
スキルスキャナースキャン .claude/commands/deploy --lenient
スキルスキャナー スキャンオール .claude/commands --recursive --lenient
# SKILL.md の代わりにカスタムメタデータファイル名を使用します
skill-scanner scan /path/to/skill --skill-file README.md
# CI/CD: 脅威が見つかった場合はビルドに失敗します
skill-scanner scan-all ./skills --fail-on-severity high --format sarif --output results.sarif
# 攻撃相関グループを含むインタラクティブな HTML レポートを生成する
skill-scanner scan /path/to/skill --use-llm --enable-meta --format html --output report.html
# カスタム YARA ルールを使用する
skill-scanner scan /path/to/skill --custom-rules /path/to/my-rules/
# カスタム分類 + 脅威マッピング プロファイル (JSON/YAML) を使用する
skill-scanner scan /path/to/skill --taxonomy /path/to/taxonomy.json --threat-mapping /path/to/threat_mapping.json
# VirusTotal ハッシュ スキャン (オプションで不明なファイルのアップロードを行う)
スキルスキャナスキャン /path/to/skill --
[切り捨てられた]
skill_scanner から SkillScanner をインポート
skill_scanner から。コア。アナライザー インポート BehavioralAnalyzer
# アナライザーを使用してスキャナーを作成する
スキャナ = SkillScanner ( アナライザー = [
ベー

avioralAnalyzer ()、
])
# スキルをスキャンする
結果 = スキャナー。 scan_skill ( "/パス/への/スキル" )
print ( f"所見: { len ( 結果 . 所見 ) } " )
print ( f"最大重大度: { result . max_severity } " )
# 注: is_safe は、HIGH/CRITICAL の所見が検出されなかったことを示します。
# スキルにリスクがまったくないことを保証するものではありません。
そうでない場合は結果を返します。 is_safe :
print ( "問題が検出されました -- 導入前に調査結果を確認してください" )
セキュリティアナライザー
アナライザー
検出方法
範囲
要件
静的
YAML + YARA パターン
すべてのファイル
なし
バイトコード
.pyc の整合性検証
Python バイトコード
なし
パイプライン
コマンド汚染分析
シェルパイプライン
なし
行動的
AST データフロー分析
Python ファイル
なし
LLM
意味解析
SKILL.md + スクリプト
APIキー
メタ
偽陽性フィルタリング
すべての調査結果
APIキー
ウイルス合計
ハッシュベースのマルウェア
バイナリファイル
APIキー
AI防御
クラウドベースのAI
テキストの内容
APIキー
CLI オプション
オプション
説明
--ポリシー
スキャン ポリシー: プリセット名 ( strict 、 Balanced 、 Permissive ) またはカスタム YAML へのパス
--使用行動
動作アナライザー (データフロー分析) を有効にする
--use-llm
LLM アナライザーを有効にする (API キーが必要)
--llm プロバイダー
CLI ルーティング用の LLM プロバイダー: anthropic または openai
--llm-consensus-runs N
LLM 分析を N 回実行し、多数派が合意した結果を維持する
--llm-max-tokens N
LLM 応答の最大出力トークン (デフォルト: 8192)
--use-virustotal
VirusTotal バイナリ スキャナーを有効にする
--vt-api-key キー
VirusTotal API キーを直接提供します (オプション)
--vt-アップロードファイル
不明なバイナリを VirusTotal にアップロードする (オプション)
--use-aidefense
Cisco AI Defense アナライザーを有効にする
--aidefense-api-url URL
AI Defense API URL をオーバーライドする (オプション)
--使用トリガー
トリガー特異性アナライザーを有効にする
--enable-meta
メタアナライザーによる誤検知フィルターの有効化
--冗長
検出結果ごとのポリシーのフィンガープリントを含め、共起します

メタデータをレンスし、メタアナライザーの誤検知を維持する
--フォーマット
出力: summary 、 json 、 markdown 、 table 、 sarif 、 html 。 HTML 形式は、折りたたみ可能な相関グループ、展開可能なコード スニペット、パイプライン テイント フロー図を含む自己完結型の対話型レポートを生成します。
--詳細
詳細な結果をマークダウン出力に含める
--コンパクト
コンパクトな JSON 出力
--出力パス
デフォルトの出力ファイルのパス ( --output-<fmt> によって上書きされます)
--失敗時の検出結果
HIGH/CRITICAL が見つかった場合はエラーで終了します ( --fail-on-severity high の短縮形)。
--fail-on-severity LEVEL
LEVEL 以上の検出結果が存在する場合はエラーで終了します (重大、高、中、低、情報)。
--カスタムルールのパス
ディレクトリからカスタム YARA ルールを使用する
--分類パス
この実行のカスタム分類プロファイル (JSON/YAML) をロードします
--threat-mapping PATH
この実行用にカスタム スキャナー脅威マッピング プロファイル (JSON) をロードします
--寛大な
不正なスキル (不正なフィールドを強制したり、デフォルト値を入力したり) を失敗するのではなく許容します。 SKILL.md が存在しない場合は、ディレクトリ内の .md ファイルのスキャンに戻ります。
--skill-file ファイル名
代わりに使用するカスタム メタデータ ファイル名
[切り捨てられた]
$ skill-scanner scan ./my-skill --use-behavioral
======================================================
スキル

[切り捨てられた]

## Original Extract

Security Scanner for Agent Skills. Contribute to cisco-ai-defense/skill-scanner development by creating an account on GitHub.

GitHub - cisco-ai-defense/skill-scanner: Security Scanner for Agent Skills · GitHub
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
cisco-ai-defense
/
skill-scanner
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
cisco-ai-defense/skill-scanner
main Branches Tags Go to file Code Open more actions menu Folders and files
61 Commits 61 Commits .cursor/ rules .cursor/ rules .devin .devin .github .github .windsurf/ rules .windsurf/ rules Formula Formula docs docs evals evals examples examples scripts scripts skill_scanner skill_scanner tests tests .env.example .env.example .gitignore .gitignore .gitleaksignore .gitleaksignore .pre-commit-config.yaml .pre-commit-config.yaml .pre-commit-hooks.yaml .pre-commit-hooks.yaml CODEOWNERS CODEOWNERS CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md FEATURE.md FEATURE.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md TESTING.md TESTING.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
A best-effort security scanner for AI Agent Skills that detects prompt injection, data exfiltration, and malicious code patterns. Combines pattern-based detection (YAML + YARA), LLM-as-a-judge , and behavioral dataflow analysis to maximize detection coverage of probable threats while minimizing false positives.
Important: This scanner provides best-effort detection, not comprehensive or complete coverage. A scan that returns no findings does not guarantee that a skill is free of all threats. See Scope and Limitations below.
Supports OpenAI Codex Skills and Cursor Agent Skills formats following the Agent Skills specification . With --lenient , also scans non-standard formats such as Claude Code .claude/commands/*.md and flat markdown skill repos.
Multi-Engine Detection - Static analysis, behavioral dataflow, LLM semantic analysis, and cloud-based scanning for layered, best-effort coverage
False Positive Filtering - Meta-analyzer significantly reduces noise while preserving detection capability
CI/CD Ready - SARIF output for GitHub Code Scanning, reusable GitHub Actions workflow , exit codes for build failures
Pre-commit Hook - Standard pre-commit framework integration to scan skills before every commit
Extensible - Plugin architecture for custom analyzers
Join the Cisco AI Discord to discuss, share feedback, or connect with the team.
Skill Scanner is a detection tool. It identifies known and probable risk patterns, but it does not certify security.
No findings ≠ no risk. A scan that returns "No findings" indicates that no known threat patterns were detected. It does not guarantee that a skill is secure, benign, or free of vulnerabilities.
Coverage is inherently incomplete. The scanner combines signature-based detection, LLM-based semantic analysis, behavioral dataflow analysis, optional cloud services, and configurable rule packs. While this approach improve coverage, no automated tool can detect every technique, especially novel or zero-day attacks.
False positives and false negatives can occur. Consensus modes and meta-analysis reduce noise, but no configuration eliminates all incorrect classifications. Tune the scan policy to your risk tolerance.
Human review remains essential. Automated scanning is one component of a defense-in-depth strategy. High-risk or production deployments should pair scanner results with manual code review and/or threat modeling.
Guide
Description
Quick Start
Get started in 5 minutes
Architecture
System design and components
Threat Taxonomy
Complete AITech threat taxonomy with examples
LLM Analyzer
LLM configuration and usage
Meta-Analyzer
False positive filtering and prioritization
Behavioral Analyzer
Dataflow analysis details
Scan Policy
Custom policies, presets, and tuning guide
Policy Quick Reference
Compact reference for policy sections and knobs
Rule Authoring
How to add signature, YARA, and Python rules
GitHub Actions
Reusable workflow for CI/CD integration
API Reference
REST API documentation
Development Guide
Contributing and development setup
Installation
Prerequisites: Python 3.10+ and uv (recommended) or pip
# Using uv (recommended)
uv pip install cisco-ai-skill-scanner
# Using pip
pip install cisco-ai-skill-scanner
Cloud Provider Extras
# AWS Bedrock support
pip install cisco-ai-skill-scanner[bedrock]
# Google AI Studio / Gemini support
pip install cisco-ai-skill-scanner[google]
# Google Vertex AI support
pip install cisco-ai-skill-scanner[vertex]
# Azure OpenAI support
pip install cisco-ai-skill-scanner[azure]
# All cloud providers
pip install cisco-ai-skill-scanner[all]
Quick Start
# For LLM analyzer and Meta-analyzer
export SKILL_SCANNER_LLM_API_KEY= " your_api_key "
export SKILL_SCANNER_LLM_MODEL= " claude-3-5-sonnet-20241022 "
# For VirusTotal binary scanning
export VIRUSTOTAL_API_KEY= " your_virustotal_api_key "
# For Cisco AI Defense
export AI_DEFENSE_API_KEY= " your_aidefense_api_key "
Interactive Wizard
Not sure which flags to use? Run skill-scanner with no arguments to launch the interactive wizard:
skill-scanner
The wizard walks you through selecting a scan target, analyzers, policy, and output format, then shows the assembled command before running it. Great for learning the CLI.
# Scan a single skill (core analyzers: static + bytecode + pipeline)
skill-scanner scan /path/to/skill
# Scan with behavioral analyzer (dataflow analysis)
skill-scanner scan /path/to/skill --use-behavioral
# Scan with all engines
skill-scanner scan /path/to/skill --use-behavioral --use-llm --use-aidefense
# Scan with meta-analyzer for false positive filtering
skill-scanner scan /path/to/skill --use-llm --enable-meta
# Scan with trigger analyzer for vague description checks
skill-scanner scan /path/to/skill --use-trigger
# Run LLM analyzer multiple times and keep majority-agreed findings
skill-scanner scan /path/to/skill --use-llm --llm-consensus-runs 3
# Scan multiple skills recursively
skill-scanner scan-all /path/to/skills --recursive --use-behavioral
# Scan multiple skills with cross-skill overlap detection
skill-scanner scan-all /path/to/skills --recursive --check-overlap
# Lenient mode: tolerate malformed skills instead of failing
skill-scanner scan /path/to/skill --lenient
skill-scanner scan-all /path/to/skills --recursive --lenient
# Lenient mode with non-standard skill formats (no SKILL.md required)
skill-scanner scan .claude/commands/deploy --lenient
skill-scanner scan-all .claude/commands --recursive --lenient
# Use a custom metadata filename instead of SKILL.md
skill-scanner scan /path/to/skill --skill-file README.md
# CI/CD: Fail build if threats found
skill-scanner scan-all ./skills --fail-on-severity high --format sarif --output results.sarif
# Generate interactive HTML report with attack correlation groups
skill-scanner scan /path/to/skill --use-llm --enable-meta --format html --output report.html
# Use custom YARA rules
skill-scanner scan /path/to/skill --custom-rules /path/to/my-rules/
# Use custom taxonomy + threat mapping profiles (JSON/YAML)
skill-scanner scan /path/to/skill --taxonomy /path/to/taxonomy.json --threat-mapping /path/to/threat_mapping.json
# VirusTotal hash scan with optional unknown-file uploads
skill-scanner scan /path/to/skill --
[truncated]
from skill_scanner import SkillScanner
from skill_scanner . core . analyzers import BehavioralAnalyzer
# Create scanner with analyzers
scanner = SkillScanner ( analyzers = [
BehavioralAnalyzer (),
])
# Scan a skill
result = scanner . scan_skill ( "/path/to/skill" )
print ( f"Findings: { len ( result . findings ) } " )
print ( f"Max severity: { result . max_severity } " )
# Note: is_safe indicates no HIGH/CRITICAL findings were detected.
# It does not guarantee the skill is free of all risk.
if not result . is_safe :
print ( "Issues detected -- review findings before deployment" )
Security Analyzers
Analyzer
Detection Method
Scope
Requirements
Static
YAML + YARA patterns
All files
None
Bytecode
.pyc integrity verification
Python bytecode
None
Pipeline
Command taint analysis
Shell pipelines
None
Behavioral
AST dataflow analysis
Python files
None
LLM
Semantic analysis
SKILL.md + scripts
API key
Meta
False positive filtering
All findings
API key
VirusTotal
Hash-based malware
Binary files
API key
AI Defense
Cloud-based AI
Text content
API key
CLI Options
Option
Description
--policy
Scan policy: preset name ( strict , balanced , permissive ) or path to custom YAML
--use-behavioral
Enable behavioral analyzer (dataflow analysis)
--use-llm
Enable LLM analyzer (requires API key)
--llm-provider
LLM provider for CLI routing: anthropic or openai
--llm-consensus-runs N
Run LLM analysis N times and keep majority-agreed findings
--llm-max-tokens N
Maximum output tokens for LLM responses (default: 8192)
--use-virustotal
Enable VirusTotal binary scanner
--vt-api-key KEY
Provide VirusTotal API key directly (optional)
--vt-upload-files
Upload unknown binaries to VirusTotal (optional)
--use-aidefense
Enable Cisco AI Defense analyzer
--aidefense-api-url URL
Override AI Defense API URL (optional)
--use-trigger
Enable trigger specificity analyzer
--enable-meta
Enable meta-analyzer for false positive filtering
--verbose
Include per-finding policy fingerprints, co-occurrence metadata, and keep meta-analyzer false positives
--format
Output: summary , json , markdown , table , sarif , html . The html format produces a self-contained interactive report with collapsible correlation groups, expandable code snippets, and pipeline taint flow diagrams
--detailed
Include detailed findings in Markdown output
--compact
Compact JSON output
--output PATH
Default output file path (overridden by --output-<fmt> )
--fail-on-findings
Exit with error if HIGH/CRITICAL found (shorthand for --fail-on-severity high )
--fail-on-severity LEVEL
Exit with error if findings at or above LEVEL exist (critical, high, medium, low, info)
--custom-rules PATH
Use custom YARA rules from directory
--taxonomy PATH
Load custom taxonomy profile (JSON/YAML) for this run
--threat-mapping PATH
Load custom scanner threat mapping profile (JSON) for this run
--lenient
Tolerate malformed skills (coerce bad fields, fill defaults) instead of failing. When SKILL.md is absent, falls back to scanning .md files in the directory
--skill-file FILENAME
Custom metadata filename to use instead
[truncated]
$ skill-scanner scan ./my-skill --use-behavioral
============================================================
Skill

[truncated]
