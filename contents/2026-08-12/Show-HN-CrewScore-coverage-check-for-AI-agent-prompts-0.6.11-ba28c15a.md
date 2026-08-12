---
source: "https://crewscore.ai/"
hn_url: "https://news.ycombinator.com/item?id=49274808"
title: "Show HN: CrewScore – coverage check for AI agent prompts (0.6.11)"
article_title: "CrewScore — find missing written guardrails"
author: "shmindmaster"
captured_at: "2026-08-12T16:46:50Z"
capture_tool: "hn-digest"
hn_id: 49274808
score: 1
comments: 0
posted_at: "2026-08-12T16:17:31Z"
tags:
  - hacker-news
  - translated
---

# Show HN: CrewScore – coverage check for AI agent prompts (0.6.11)

- HN: [49274808](https://news.ycombinator.com/item?id=49274808)
- Source: [crewscore.ai](https://crewscore.ai/)
- Score: 1
- Comments: 0
- Posted: 2026-08-12T16:17:31Z

## Translation

タイトル: HN を表示: CrewScore – AI エージェント プロンプトのカバレッジ チェック (0.6.11)
記事のタイトル: CrewScore — 記載されていないガードレールを見つける
説明: 23 のパブリックなブラウザー ローカル コントロールを使用して、AI アシスタントの指示に欠けている書かれたガードレールを見つけます。

記事本文:
CrewScore — 欠けている書かれたガードレールを見つける
クルースコア
製品例 ルール CI 方法論
フィードバックを送信する
GitHub
開発者モード
ブラウザローカルのガードレールのチェックリスト
AI エージェントのプロンプトが忘れた安全ルールを見つけます。
AI の指示を貼り付けます。 CrewScore は、ブラウザー内の 23 のパブリック ガードレール コントロールに対してそれらをチェックします。プロンプトのアップロード、サインアップ、API キー、モデル呼び出しは行われません。
書面による管理の対象範囲であり、実行時の証明ではありません。エージェントがガードレールに従っているかどうかではなく、ガードレールが書き留められているかどうかをチェックします。
最初のギャップ ローカルチェック待ち
合成例はローカルでチェックされます。 CrewScore は書面による管理の範囲と最初に確認する管理を報告します。実行時の動作はテストされません。
83 の実稼働ラベル付きエージェント プロンプトのカバレッジ中央値 10/100
356 個のプロンプトをオフラインでスコア付けしました。品質ランキングではありません。メソッドとハーネスです。
AI 命令が存在する場所を選択するか、下に貼り付けます。
ChatGPT カスタム命令、クロード プロジェクト プロンプト、サポート ボット プロンプト、または別のシステム プロンプトを貼り付けます。
サポートされている: public github.com/.../blob/... および raw.githubusercontent.com ファイルのみ。
デモから始めるか、手順を貼り付けてください。 CrewScore は、検出した書き込みコントロールと、レビューするよう提案されたコントロールを表示します。
このプロンプトに含まれる文言のみを選択してください。これらは編集可能なテキストの提案であり、実行時の強制ではありません。
作業コピーに適用 レビューをキャンセル 書かれたガードレールと実行時の承認、ツールの制限、および評価をペアにします。 「ランタイムの次のステップ」を参照してください。
手順を貼り付け、記載されていないガードレールを見つけて、文言を改善するか、定期的なリポジトリ チェックを追加します。
すべてのコントロール、正規表現、ルールセットのバージョンは公開されています。開発者モードでは、完全一致と ID が表示されます。
CLI またはアクションを使用して、特定の記述されたコントロールを保護し、プロンプトなしの SARIF 結果を生成します。
カバレッジは ID に役立ちます

省略された書き込みコントロールを確認します。これは、品質ランキング、認定、または実行時テストではありません。
いいえ。採点はローカルで行われます。プロンプト テキストがこのブラウザから離れることはありません。
いいえ。これは、書かれたコントロールのパブリック テキスト パターンが検出されたことを意味します。実行時の強制とテストは分離されたままになります。
はい。オープンソースの CLI または GitHub Action を使用して、特定のコントロールを要求し、プロンプトのないベースラインを保護し、SARIF を発行します。
採点はローカルで行われます。プロンプト テキストがブラウザから離れることはありません。匿名の許可リストに登録された使用状況イベントには、プロンプト テキストが含まれません。プライバシーの詳細をお読みください。
Sarosh Hussain によって作成および保守されています。 Pendoah は CrewScore のコンテキストを運営する会社です。技術的主張はリポジトリに文書化され、検証資料が引用されます。
このデバイスの匿名使用イベントを無効にします。
pip install crewscore · 使用: shmindmaster/crewscore@v2 · セキュリティ · ベンダー チェックリスト · フィードバックを送信 · GitHub ·
ChatGPT: 設定 → パーソナライゼーション → カスタム指示。カスタム手順のテキストをコピーします。
クロード: プロジェクト → プロジェクト指示 (または Claude.ai カスタム指示) を開きます。指示文をコピーします。
カーソル: リポジトリのルートで AGENTS.md 、 .cursorrules 、またはプロジェクト ルールを開きます。コーディング エージェント ファイルは、ガバナンス グレードではなく、構成の匂いに関してスコア付けされます。
チームまたはアプリ: システム プロンプト、アシスタント セットアップ、または AI 指示フィールドを探します。検査が許可されているテキストのみをコピーしてください。

## Original Extract

Find written guardrails missing from AI assistant instructions with 23 public, browser-local controls.

CrewScore — find missing written guardrails
CrewScore
Product Examples Rules CI Methodology
Give feedback
GitHub
Developer mode
Browser-local guardrail checklist
Find the safety rules your AI agent prompt forgot.
Paste your AI instructions. CrewScore checks them against 23 public guardrail controls in your browser — no prompt upload, signup, API key, or model call.
Written-control coverage, not runtime proof. It checks whether guardrails are written down — not whether an agent follows them.
First gap Waiting for local check
A synthetic example is checked locally. CrewScore reports written-control coverage and the first control to review; it does not test runtime behavior.
10/100 median coverage among 83 production-labeled agent prompts
356 prompts scored offline · not a quality ranking · method & harness
Pick where your AI instructions live — or paste them below.
Paste ChatGPT custom instructions, a Claude Project prompt, a support-bot prompt, or another system prompt.
Supported: public github.com/.../blob/... and raw.githubusercontent.com files only.
Start with a demo or paste instructions. CrewScore will show written controls it detected and suggested controls to review.
Select only wording that belongs in this prompt. These are editable text suggestions, not runtime enforcement.
Apply to working copy Cancel review Pair written guardrails with runtime approvals, tool restrictions, and evaluation. See runtime next steps .
Paste instructions, find missing written guardrails, then improve the wording or add a recurring repository check.
Every control, regex, and ruleset version is public. Developer mode shows exact matches and IDs.
Use the CLI or Action to protect specific written controls and generate prompt-free SARIF findings.
Coverage helps identify omitted written controls. It is not a quality ranking, a certification, or a runtime test.
No. Scoring happens locally; prompt text never leaves this browser.
No. It means public text patterns for written controls were detected. Runtime enforcement and testing remain separate.
Yes. Use the open-source CLI or GitHub Action to require particular controls, protect a prompt-free baseline, and emit SARIF.
Scoring happens locally. Your prompt text never leaves your browser. Anonymous allowlisted usage events exclude prompt text; read privacy details .
Created and maintained by Sarosh Hussain . Pendoah is the company operating context for CrewScore; technical claims are documented in the repository and cited validation material.
Use disable anonymous usage events on this device .
pip install crewscore · uses: shmindmaster/crewscore@v2 · Security · Vendor checklist · Give feedback · GitHub ·
ChatGPT: Settings → Personalization → Custom Instructions. Copy the custom instructions text.
Claude: open a Project → Project instructions (or Claude.ai custom instructions). Copy the instruction text.
Cursor: open AGENTS.md , .cursorrules , or project rules in the repo root. Coding-agent files are scored for config smells, not a governance grade.
Teams or apps: look for a system prompt, assistant setup, or AI instructions field. Copy only text you are allowed to inspect.
