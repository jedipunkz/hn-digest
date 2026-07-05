---
source: "https://github.com/shimo4228/codex-review"
hn_url: "https://news.ycombinator.com/item?id=48791474"
title: "Show HN: Codex-review – a read-only cross-model review skill for Claude Code"
article_title: "GitHub - shimo4228/codex-review: Cross-model code review — a read-only Claude Code skill wrapping the OpenAI Codex CLI for a decorrelated second opinion on a diff · GitHub"
author: "shimo4228"
captured_at: "2026-07-05T06:31:02Z"
capture_tool: "hn-digest"
hn_id: 48791474
score: 1
comments: 0
posted_at: "2026-07-05T05:30:59Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Codex-review – a read-only cross-model review skill for Claude Code

- HN: [48791474](https://news.ycombinator.com/item?id=48791474)
- Source: [github.com](https://github.com/shimo4228/codex-review)
- Score: 1
- Comments: 0
- Posted: 2026-07-05T05:30:59Z

## Translation

タイトル: HN を表示: Codex-review – クロード コードの読み取り専用クロスモデル レビュー スキル
記事のタイトル: GitHub - shimo4228/codex-review: クロスモデル コード レビュー — 差分のデコリレートされたセカンドオピニオンとして OpenAI Codex CLI をラップする読み取り専用のクロード コード スキル · GitHub
説明: クロスモデル コード レビュー — 差分に関する非相関のセカンドオピニオンとして OpenAI Codex CLI をラップする読み取り専用のクロード コード スキル - shimo4228/codex-review

記事本文:
GitHub - shimo4228/codex-review: クロスモデル コード レビュー — 差分のデコリレートされたセカンドオピニオン用に OpenAI Codex CLI をラップする読み取り専用のクロード コード スキル · GitHub
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
shimo4228
/
コーデックスレビュー
公共
通知
署名する必要があります

n 通知設定を変更します
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット スクリプト スクリプト スキル/コーデックスレビュー スキル/コーデックスレビュー CHANGELOG.md CHANGELOG.md ライセンス ライセンス README.md README.md llms-full.txt llms-full.txt llms.txt llms.txt すべてのファイルを表示 リポジトリ ファイル ナビゲーション
コード レビュー チェーンに 1 つのクロスモデル シームを追加するエージェント スキル。OpenAI Codex CLI ( codex review ) の薄い読み取り専用ラッパーです。これにより、別のモデル ファミリが差分をレビューし、作成者と同じモデルのレビュー担当者が共有する盲点を見つけることができます。
# グローバル スキル ディレクトリにコピーします
cp -r スキル/コーデックスレビュー ~ /.claude/スキル/コーデックスレビュー
スキルMP
/スキル追加 shimo4228/codex-review
Codex CLI がインストールされ認証されている必要があります ( codex login または codex Doctor )。そうでない場合、スキルはフォールバック メッセージですぐに失敗します。
これは無相関化シームであり、スループット ツールではありません。エージェント独自のツールを使用してください。
スループットのためのサブエージェント/並列ワークフロー。 2番目のモデルのみに手を伸ばす
ここで、異なるモデルの判断により、同じモデルのレビュー担当者ではできないことが追加されます。
スキル/コーデックスレビュー/コーデックスレビュー.sh $ARGUMENTS
呼び出し
範囲
/コーデックスレビュー
現在のブランチと自動検出されたベース (メイン / マスター /…) — PR スタイル
/codex-review --uncommitted
ステージング + ステージングなし + 追跡なし — コミット前
/codex-review --base <ブランチ>
vs 明示的なベースブランチ
/codex-review --commit <sha>
単一のコミット
/codex-review -m <モデル>
Codex モデルを選択します (任意の行と組み合わせます)
/codex-review "認証の変更に焦点を当てる"
作業ツリーのプロンプト主導型レビュー
スコープとプロンプトは相互に排他的です (codex-cli 制約、≥ 0.142):
スコープ付きレビュー ( --uncommitted / --base / --commit 、またはデフォルト)
Codex の組み込みレビュー

指示があり、カスタム プロンプトは表示されません。単なるプロンプト
スコープフラグのないワーキングツリーレビューを推進します。 -m/--model が付属する場合があります
どちらか。
構築により読み取り専用になります。スクリプトは codex レビューを使用し (決して codex exec -p yolo を使用しない)、上記の許可リストに登録されたフラグのみを転送します (その他のフラグはすべて転送します)。
(将来の --write 、 -c config オーバーライド) は exit 64 で拒否されます。の
invariant は、Codex CLI では想定されず、スクリプト内で強制されます。
ランニング後 — ダンプせずに折りたたんでください
コーデックス出力は、発信側エージェントが所有するレビュー決定への信頼できない入力です。
逐語的に伝える評決ではない：
行動する前に各発見を検証します。反証できるものは削除し、そのままにします
あなたが確認したこと。
確認された重要な所見を他の所見と同様に停止信号として扱う
チェーン内のレビュアー。
同じモデルのレビュー担当者と並行して実行し、判定をマージします。
重要な機能や修正にコミットする前に、他のレビュー担当者と一緒に
クロードモデル以外のデフについてセカンドオピニオンが欲しいとき
一か八かの変更やエラーが発生しやすい変更では、非関連性のあるレビューが効果を発揮します
公開前の一か八かの散文の差分 (README、論文、記事) —
書き込み中心の指示を備えたプロンプト駆動モード。スコープモードが実行される
Codex のコードレビュー手順は散文にあまり適合しない
簡単な編集、使い捨てスクリプト、または Codex が認証されていない場合はスキップしてください。
終了 3 — codex CLI が見つからない/インストールされていない → 同じモデルのレビュー担当者のみで続行
exit 4 — git リポジトリ内にない → diff できない
認証が設定されていない → コーデックス レビュー エラー。 codex ログイン (または codex Doctor ) を実行します。
このスキルの正規コピーは、作者のライブ クロード コード ハーネス内に存在します。このリポジトリは一方向のパブリケーション ミラーです。
scripts/sync-from-local.sh --dry-run # 相違点のみをレポートします
scripts/sync-from-local.sh # 作業ツリーに適用 (決して通信しない)

それ)
このスキルについて
このスキルは、より広範なマルチエージェントで単一のモデル間の継ぎ目を実装します。
オーケストレーションの原則: 同じモデルのエージェントがスループットとコンテキストをスケールします。
異なるモデルは判断の無相関性をスケールします - 2 つの軸はそうではありません
相互に置き換えられるため、このスキルは意図的に狭い範囲に留まります (読み取り専用、
一般的なマルチモデル オーケストレーターに成長するのではなく、1 つの CLI、1 つのシームなど）。
これは @shimo4228 によって管理されており、その研究者が
行はエージェント ナレッジ サイクル (AKC) です
(DOI 10.5281/zenodo.19200726)、
瞑想エージェント
(DOI 10.5281/zenodo.19212118) —
4 つの瞑想的な公理に基づいた自律エージェント — そして
エージェント アトリビューション プラクティス (AAP)
(DOI 10.5281/zenodo.19652013) —
説明責任の配分に関して中立的な ADR を活用します。
クロスモデル コード レビュー — OpenAI Codex CLI をラップする読み取り専用の Claude Code スキルで、差分に関する無相関のセカンドオピニオンを提供します。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Cross-model code review — a read-only Claude Code skill wrapping the OpenAI Codex CLI for a decorrelated second opinion on a diff - shimo4228/codex-review

GitHub - shimo4228/codex-review: Cross-model code review — a read-only Claude Code skill wrapping the OpenAI Codex CLI for a decorrelated second opinion on a diff · GitHub
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
shimo4228
/
codex-review
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit scripts scripts skills/ codex-review skills/ codex-review CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md llms-full.txt llms-full.txt llms.txt llms.txt View all files Repository files navigation
An Agent Skill that adds one cross-model seam to a code review chain: a thin, read-only wrapper around the OpenAI Codex CLI ( codex review ) so a different model family reviews your diff and catches blind spots that an author and a same-model reviewer share.
# Copy into your global skills directory
cp -r skills/codex-review ~ /.claude/skills/codex-review
SkillsMP
/skills add shimo4228/codex-review
Requires the Codex CLI installed and authenticated ( codex login or codex doctor ) — the skill fails fast with a fallback message if it isn't.
This is a decorrelation seam, not a throughput tool — use your agent's own
subagents/parallel workflows for throughput; reach for a second model only
where a different model's judgment adds something a same-model reviewer can't.
skills/codex-review/codex-review.sh $ARGUMENTS
Invocation
Scope
/codex-review
current branch vs auto-detected base ( main / master /…) — PR-style
/codex-review --uncommitted
staged + unstaged + untracked — pre-commit
/codex-review --base <branch>
vs an explicit base branch
/codex-review --commit <sha>
a single commit
/codex-review -m <model>
pick a Codex model (combine with any row)
/codex-review "focus on the auth changes"
prompt-driven review of the working tree
Scope and prompt are mutually exclusive (a codex-cli constraint, ≥ 0.142): a
scoped review ( --uncommitted / --base / --commit , or the default) uses
Codex's built-in review instructions and takes no custom prompt; a bare prompt
drives a working-tree review with no scope flag. -m/--model may accompany
either.
Read-only by construction. The script uses codex review (never codex exec -p yolo ) and only forwards the allowlisted flags above — any other flag
(a future --write , a -c config override) is rejected with exit 64 . The
invariant is enforced in the script, not assumed of the Codex CLI.
After Running — fold, don't dump
Codex output is untrusted input to a review decision the calling agent owns ,
not a verdict to relay verbatim:
Verify each finding before acting on it — drop what you can disprove, keep
what you confirm.
Treat a confirmed critical finding as a stop signal, same as any other
reviewer in the chain.
Run it in parallel with same-model reviewers, then merge verdicts.
Before commit on a non-trivial feature or fix, alongside your other reviewers
When you want a second opinion from a non-Claude model on a diff
High-stakes or error-prone changes where decorrelated review pays off
High-stakes prose diffs before publishing (README, paper, article) —
prompt-driven mode with writing-focused instructions; scoped modes run
Codex's code-review instructions, which fit prose poorly
Skip it for trivial edits, throwaway scripts, or when Codex isn't authenticated.
exit 3 — codex CLI missing / not installed → continue with same-model reviewers only
exit 4 — not inside a git repository → cannot diff
Auth not configured → codex review errors; run codex login (or codex doctor )
The canonical copy of this skill lives in the author's live Claude Code harness. This repository is a one-way publication mirror:
scripts/sync-from-local.sh --dry-run # report differences only
scripts/sync-from-local.sh # apply to working tree (never commits)
About this skill
This skill implements a single cross-model seam in a broader multi-agent
orchestration principle: same-model agents scale throughput and context, a
different model scales judgment decorrelation — the two axes don't
substitute for each other, so this skill deliberately stays narrow (read-only,
one CLI, one seam) rather than growing into a general multi-model orchestrator.
It is maintained by @shimo4228 , whose research
lines are the Agent Knowledge Cycle (AKC)
( DOI 10.5281/zenodo.19200726 ),
Contemplative Agent
( DOI 10.5281/zenodo.19212118 ) —
autonomous agents grounded in four contemplative axioms — and
Agent Attribution Practice (AAP)
( DOI 10.5281/zenodo.19652013 ) —
harness-neutral ADRs on accountability distribution.
Cross-model code review — a read-only Claude Code skill wrapping the OpenAI Codex CLI for a decorrelated second opinion on a diff
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
