---
source: "https://github.com/Higangssh/yocto-agent-skills"
hn_url: "https://news.ycombinator.com/item?id=48726866"
title: "Show HN: Yocto/BitBake skills that make AI agents check official docs"
article_title: "GitHub - Higangssh/yocto-agent-skills: Official-doc-first Yocto Project and BitBake skills for AI coding agents · GitHub"
author: "swq115"
captured_at: "2026-06-30T00:31:11Z"
capture_tool: "hn-digest"
hn_id: 48726866
score: 2
comments: 0
posted_at: "2026-06-29T23:52:46Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Yocto/BitBake skills that make AI agents check official docs

- HN: [48726866](https://news.ycombinator.com/item?id=48726866)
- Source: [github.com](https://github.com/Higangssh/yocto-agent-skills)
- Score: 2
- Comments: 0
- Posted: 2026-06-29T23:52:46Z

## Translation

タイトル: Show HN: AI エージェントに公式ドキュメントをチェックさせる Yocto/BitBake スキル
記事タイトル: GitHub - Higangssh/yocto-agent-skills: AI コーディング エージェントのための公式ドキュメントファースト Yocto プロジェクトと BitBake スキル · GitHub
説明: AI コーディング エージェント向けの公式ドキュメント初の Yocto プロジェクトと BitBake スキル - Higangssh/yocto-agent-skills

記事本文:
GitHub - Higangssh/yocto-agent-skills: AI コーディング エージェント向けの公式ドキュメントファースト Yocto プロジェクトと BitBake スキル · GitHub
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
ヒガンシュ
/
yocto-エージェント-スキル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のna

違反オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット .github .github エージェント エージェント evals evals 例 例 リファレンス リファレンス スキル スキル CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.ko.md README.ko.md README.md README.md SKILL.md SKILL.md すべてのファイルを表示リポジトリ ファイルのナビゲーション
公式ドキュメントファーストの Yocto プロジェクトと AI コーディング エージェント向けの BitBake スキル。
Yocto はリリースに敏感で、詳細に設定可能であり、一般的な LLM が幻覚を起こしやすいです。このリポジトリは、エージェントに、公式ドキュメントへのルーティング、BitBake エラーのデバッグ、レシピのレビュー、レイヤーのレビュー、イメージ/rootfs の問題の診断、BSP/カーネルの問題の処理、セキュリティ/SBOM ワークフローの処理に焦点を当てたスキルを提供します。
韓国語ドキュメント: README.ko.md
skill/yocto-doc-router : 正しい Yocto、OpenEmbedded、および BitBake の公式ドキュメントへのリリース対応ルーティング。
skill/bitbake-debug : BitBake ビルドの失敗に対する task/log/rootfs/package/provider のデバッグ。
skill/yocto-recipe-review : レシピ、bbappend、bbclass、依存関係、パッケージ化、ライセンス、およびオーバーライド構文のレビュー。
skill/yocto-layer-review :layer.conf、レイヤー互換性、優先度、依存関係、プロバイダー、および bbappend マッチング レビュー。
skill/yocto-image-rootfs : イメージ レシピ、パッケージ名、 IMAGE_INSTALL 、 IMAGE_FEATURES 、 do_rootfs 、 pkgdata 、およびパッケージ マネージャーの問題。
skill/yocto-bsp-kernel : マシン構成、BSP レイヤー、カーネル プロバイダー、devicetree、defconfig、U-Boot、およびデプロイ アーティファクト。
skill/yocto-security-sbom : ライセンス メタデータ、CVE チェック、SPDX/SBOM、アーカイバー/コピーレフト フロー、およびコンプライアンス アーティファクト。
ルート SKILL.md は、リポジトリを単一のスキルとしてインストールするホストの互換ルーターとして残ります。
デブ

BitBake タスクの失敗の記録: do_fetch 、 do_unpack 、 do_patch 、 do_configure 、 do_compile 、 do_install 、 do_package 、 do_package_qa 、 do_rootfs 、 do_image
.bb 、 .bbappend 、 .bbclass 、イメージレシピ、マシン構成、ディストリビューション構成、および Layer.conf の作成とレビュー
BitBake オーバーライド構文の最新化: VAR:append 、 FILES:${PN} 、 RDEPENDS:${PN} 、タスク オーバーライド、およびパッケージ オーバーライド
レイヤー、bbappend、プロバイダーの選択、パッケージ分割、rootfs エラー、QA メッセージ、カーネル/BSP メタデータ、およびイメージ構成のレビュー
DEPENDS と RDEPENDS 、レシピ名とパッケージ名、 SRCREV 、 LIC_FILES_CHKSUM 、 INSANE_SKIP 、ホストの汚染、および sstate のクリーンアップに関する一般的な AI の間違いを削減します。
コレクション対応エージェントの場合は、skills/ の下に個別のフォルダーをインストールします。
mkdir -p " ${CODEX_HOME :- $HOME / .codex} /skills "
yocto-doc-router bitbake-debug yocto-recipe-review yocto-layer-review yocto-image-rootfs yocto-bsp-kernel yocto-security-sbom のスキルについて。する
ln -s " $( pwd ) /skills/ $skill " " ${CODEX_HOME :- $HOME / .codex} /skills/ $skill "
完了しました
1 つのフォルダーを 1 つのスキルとしてインストールするホストの場合は、リポジトリ ルートをインストールします。
mkdir -p " ${CODEX_HOME :- $HOME / .codex} /skills "
ln -s " $( pwd ) " " ${CODEX_HOME :- $HOME / .codex} /skills/yocto-agent-skills "
プロンプトの例
bitbake-debug を使用して、この do_rootfs エラーを診断します。
yocto-recipe-review を使用してこのレシピを確認し、オーバーライド構文を最新化します。
yocto-layer-review を使用して、この bbappend が適用されない理由を説明します。
yocto-image-rootfs を使用して、パッケージが最終イメージに含まれていない理由を調べます。
yocto-bsp-kernel を使用して、デバイスツリーがデプロイ/イメージから欠落している理由をデバッグします。
yocto-security-sbom を使用して、このライセンス チェックサムと SBOM 設定を確認します。
内容
SKILL.md : ルート互換ルーター
skill/*/SKILL.md : 注目のインスタ

習得可能なスキル
References/shared/official-doc-map.md : 公式 Yocto/BitBake ドキュメントのルーティング (問題タイプ別)
References/shared/yocto-field-guide.md : レシピ、レイヤー、タスク、QA、イメージ、プロバイダー、BSP/カーネル作業のコンパクトなフィールド ガイド
References/bitbake/variables-core.md : エージェントがよく混乱する変数
References/bitbake/classes-core.md : 共通クラスとレビュールール
References/bitbake/tasks-reference.md : タスクレベルのデバッグリファレンス
References/yocto/qa-errors.md : 一般的な QA エラー パターン
References/yocto/migration.md : リリース対応の移行チェック
References/yocto/image-rootfs.md : イメージと rootfs のトラブルシューティング
References/yocto/bsp-kernel.md : BSP とカーネルのトラブルシューティング
References/yocto/security-sbom.md : セキュリティ、ライセンス、CVE、および SBOM ワークフロー
Examples/ : 現実的な失敗例と予想される回答パターン
evals/prompts.md : 手動フォワードテストプロンプト
Agents/openai.yaml : 互換性のあるスキルホストの UI メタデータ
公式ドキュメントファーストの Yocto プロジェクトと AI コーディング エージェント向けの BitBake スキル
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
v0.2.0
最新の
2026 年 6 月 29 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Official-doc-first Yocto Project and BitBake skills for AI coding agents - Higangssh/yocto-agent-skills

GitHub - Higangssh/yocto-agent-skills: Official-doc-first Yocto Project and BitBake skills for AI coding agents · GitHub
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
Higangssh
/
yocto-agent-skills
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits .github .github agents agents evals evals examples examples references references skills skills CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.ko.md README.ko.md README.md README.md SKILL.md SKILL.md View all files Repository files navigation
Official-doc-first Yocto Project and BitBake skills for AI coding agents.
Yocto is release-sensitive, deeply configurable, and easy for general LLMs to hallucinate. This repository gives agents focused skills for routing to official documentation, debugging BitBake failures, reviewing recipes, reviewing layers, diagnosing image/rootfs problems, working through BSP/kernel issues, and handling security/SBOM workflows.
Korean documentation: README.ko.md
skills/yocto-doc-router : release-aware routing to the right Yocto, OpenEmbedded, and BitBake official documentation.
skills/bitbake-debug : task/log/rootfs/package/provider debugging for BitBake build failures.
skills/yocto-recipe-review : recipe, bbappend, bbclass, dependency, packaging, licensing, and override syntax review.
skills/yocto-layer-review : layer.conf, layer compatibility, priority, dependency, provider, and bbappend matching review.
skills/yocto-image-rootfs : image recipes, package names, IMAGE_INSTALL , IMAGE_FEATURES , do_rootfs , pkgdata, and package manager issues.
skills/yocto-bsp-kernel : machine config, BSP layers, kernel providers, devicetree, defconfig, U-Boot, and deploy artifacts.
skills/yocto-security-sbom : license metadata, CVE checks, SPDX/SBOM, archiver/copyleft flows, and compliance artifacts.
The root SKILL.md remains as a compatibility router for hosts that install a repository as a single skill.
Debugging BitBake task failures: do_fetch , do_unpack , do_patch , do_configure , do_compile , do_install , do_package , do_package_qa , do_rootfs , do_image
Writing and reviewing .bb , .bbappend , .bbclass , image recipes, machine config, distro config, and layer.conf
Modernizing BitBake override syntax: VAR:append , FILES:${PN} , RDEPENDS:${PN} , task overrides, and package overrides
Reviewing layers, bbappends, provider selection, package splitting, rootfs failures, QA messages, kernel/BSP metadata, and image composition
Reducing common AI mistakes around DEPENDS vs RDEPENDS , recipe names vs package names, SRCREV , LIC_FILES_CHKSUM , INSANE_SKIP , host contamination, and sstate cleanup
For collection-aware agents, install the individual folders under skills/ .
mkdir -p " ${CODEX_HOME :- $HOME / .codex} /skills "
for skill in yocto-doc-router bitbake-debug yocto-recipe-review yocto-layer-review yocto-image-rootfs yocto-bsp-kernel yocto-security-sbom ; do
ln -s " $( pwd ) /skills/ $skill " " ${CODEX_HOME :- $HOME / .codex} /skills/ $skill "
done
For hosts that install one folder as one skill, install the repository root:
mkdir -p " ${CODEX_HOME :- $HOME / .codex} /skills "
ln -s " $( pwd ) " " ${CODEX_HOME :- $HOME / .codex} /skills/yocto-agent-skills "
Example Prompts
Use bitbake-debug to diagnose this do_rootfs failure.
Use yocto-recipe-review to review this recipe and modernize the override syntax.
Use yocto-layer-review to explain why this bbappend is not being applied.
Use yocto-image-rootfs to find why my package is not in the final image.
Use yocto-bsp-kernel to debug why my devicetree is missing from deploy/images.
Use yocto-security-sbom to review this license checksum and SBOM setup.
Contents
SKILL.md : root compatibility router
skills/*/SKILL.md : focused installable skills
references/shared/official-doc-map.md : official Yocto/BitBake documentation routing by problem type
references/shared/yocto-field-guide.md : compact field guide for recipes, layers, tasks, QA, images, providers, and BSP/kernel work
references/bitbake/variables-core.md : variables agents often confuse
references/bitbake/classes-core.md : common classes and review rules
references/bitbake/tasks-reference.md : task-level debugging reference
references/yocto/qa-errors.md : common QA error patterns
references/yocto/migration.md : release-aware migration checks
references/yocto/image-rootfs.md : image and rootfs troubleshooting
references/yocto/bsp-kernel.md : BSP and kernel troubleshooting
references/yocto/security-sbom.md : security, license, CVE, and SBOM workflows
examples/ : realistic failure examples and expected answer patterns
evals/prompts.md : manual forward-test prompts
agents/openai.yaml : UI metadata for compatible skill hosts
Official-doc-first Yocto Project and BitBake skills for AI coding agents
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
v0.2.0
Latest
Jun 29, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
