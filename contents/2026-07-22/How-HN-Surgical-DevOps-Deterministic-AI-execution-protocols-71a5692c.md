---
source: "https://github.com/bonushora/surgical-dev-ops"
hn_url: "https://news.ycombinator.com/item?id=49010491"
title: "How HN: Surgical DevOps – Deterministic AI execution protocols"
article_title: "GitHub - bonushora/surgical-dev-ops: core-methodologies · GitHub"
author: "bonus_dev"
captured_at: "2026-07-22T18:03:15Z"
capture_tool: "hn-digest"
hn_id: 49010491
score: 1
comments: 0
posted_at: "2026-07-22T17:39:22Z"
tags:
  - hacker-news
  - translated
---

# How HN: Surgical DevOps – Deterministic AI execution protocols

- HN: [49010491](https://news.ycombinator.com/item?id=49010491)
- Source: [github.com](https://github.com/bonushora/surgical-dev-ops)
- Score: 1
- Comments: 0
- Posted: 2026-07-22T17:39:22Z

## Translation

タイトル: How HN: Surgical DevOps – 決定論的 AI 実行プロトコル
記事のタイトル: GitHub - Bonushora/surgical-dev-ops: core-methodologies · GitHub
説明: コア方法論。 GitHub でアカウントを作成して、bonushora/surgical-dev-ops の開発に貢献してください。

記事本文:
GitHub - Bonushora/surgical-dev-ops: コア方法論 · GitHub
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
コードの品質 マージ時に品質を強制する
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
ボーナスホラ
/
外科開発運用
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
タラ

そして
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
38 コミット 38 コミット docs docs プロトコル プロトコル APPLICABILITY_EN.md APPLICABILITY_EN.md CHANGELOG.md CHANGELOG.md COTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md README_EN.md README_EN.md ROADMAP.md ROADMAP.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Surgical DevOps は、AI 支援ソフトウェア開発ライフサイクル中の大規模言語モデル (LLM) の動作を管理および標準化するために作成された、オープンソースのプロトコルに依存しないエコシステムです。
その目標は、回帰を減らし、既存のシステムについての過度の仮定を回避し、長期の開発セッション中に戦略的知識を保持することです。
エコシステムは、次の 2 つの主要なプロトコルの組み合わせを通じて動作します。
BH-SEP (Safe Evolution Protocol): AI が既存のソフトウェアを安全に進化させる方法を定義します。このプロトコルは、変更前の検査 (Inspect First)、機能コードの保存 (Preserve Everything)、最小限の変更 (Minimal Diff)、および増分検証を確立します。
BH-SDP (スナップショット & デリバリー プロトコル): 状態の保存とセッション間の継続性のメカニズムを定義します。このプロトコルは、意思決定、契約、関連するファイル、および次のステップを、コンテキストを安全に転送できる構造化された成果物 (スナップショット) に変換します。
従来のモデル (回帰への道)
[プロンプト] → [AI 精神的再構築] → [既存のコード書き換え] → [バグ/リグレッション]
外科的 DevOps モデル (安全かつ持続的な進化)
[既存のコード (True)] ──> [完全検査] ──> [最小差分] ──> [検証] ──> [ステータス スナップショット]

o] ──> [安全な次のステップ]
🏛️ エコシステムの基本原則
最初に検査します:
既存のコードは信頼できる情報源を表します。 AI は、適切な検査なしに契約、ルート、依存関係、または状態管理を引き継いではいけません。
すべてを保存:
機能コードは保存する必要があります。要求された範囲外の変更はリスクを高めるため、避けるべきです。
最小の差分:
進化は、プロジェクトの歴史への影響を軽減する、小規模で個別の追跡可能な介入を通じて発生する必要があります。
すぐに検証します:
次のステップに進む前に、各変更を検証する必要があります。
段階的に前進:
複雑な問題は、より小さな独立したステップに分割する必要があります。
状態の継続性:
セッション間の安全な継続を可能にするために、関連する決定、契約、および停止ポイントを保存する必要があります。
🤖 アーティファクト: AI 用の統合システム プロンプト
Surgical DevOps エコシステムを使用して開発セッションを開始するには、以下をコピーして貼り付けます。
https://raw.githubusercontent.com/bonushora/surgical-dev-ops/main/protocols/BH-SEP.md
https://raw.githubusercontent.com/bonushora/surgical-dev-ops/main/protocols/BH-SDP.md
BH-SEP (Safe Evolution Protocol) と BH-SDP (Snapshot & Delivery Protocol) のガイドラインを厳密に組み合わせて黙って採用します。
プロジェクト エコシステムを専門とするシニア ソフトウェア エンジニアとして活動します。
変更を加える前に、既存のコンテキストを検査してください。動作中のコードを保持し、最小限の変更を適用して、各ステップを検証します。
以前のセッションのスナップショットが存在する場合は、それを使用して現在の状態を回復します。
後

プロトコルと最初のコンテキストを理解したら、次のように答えて確認します。
次に、最初に検査する必要があるファイルまたはコンテキストを要求します。
BH-SEP — 安全進化プロトコル
BH-SDP — スナップショットおよび配信プロトコル
Surgical DevOps は BônusHora エコシステム内で生まれましたが、その原則は言語、フレームワーク、アーキテクチャから独立しています。
AI 支援を利用して開発されたプロジェクト。
外科的 DevOps はエンジニアリング上の判断に代わるものではありません。
これにより、人間の決定が主権を維持し、AI 支援による実行が整合性、追跡可能性、安全性を維持する、規律あるモデルが確立されます。
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

core-methodologies. Contribute to bonushora/surgical-dev-ops development by creating an account on GitHub.

GitHub - bonushora/surgical-dev-ops: core-methodologies · GitHub
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
Code Quality Enforce quality at merge
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
bonushora
/
surgical-dev-ops
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
38 Commits 38 Commits docs docs protocols protocols APPLICABILITY_EN.md APPLICABILITY_EN.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md README_EN.md README_EN.md ROADMAP.md ROADMAP.md View all files Repository files navigation
O Surgical DevOps é um ecossistema protocolar agnóstico e de código aberto criado para governar e padronizar o comportamento de Grandes Modelos de Linguagem (LLMs) durante o ciclo de desenvolvimento de software assistido por Inteligência Artificial.
Seu objetivo é reduzir regressões, evitar suposições indevidas sobre sistemas existentes e preservar conhecimento estratégico durante sessões longas de desenvolvimento.
O ecossistema opera através da combinação de dois protocolos principais:
BH-SEP (Safe Evolution Protocol): Define como a IA deve evoluir software existente com segurança. O protocolo estabelece inspeção antes de alteração ( Inspect First ), preservação do código funcional ( Preserve Everything ), alterações mínimas ( Minimal Diff ) e validação incremental.
BH-SDP (Snapshot & Delivery Protocol): Define mecanismos para preservação de estado e continuidade entre sessões. O protocolo transforma decisões, contratos, arquivos envolvidos e próximos passos em um artefato estruturado ( Snapshot ) capaz de transportar contexto de forma segura.
O Modelo Tradicional (Caminho para Regressões)
[Prompt] ──> [Reconstrução Mental da IA] ──> [Reescrita de Código Existente] ──> [Bug / Regressão]
O Modelo Surgical DevOps (Evolução Segura e Persistente)
[Código Existente (Verdade)] ──> [Inspeção Completa] ──> [Diff Mínimo] ──> [Validação] ──> [Snapshot de Estado] ──> [Próximo Passo Seguro]
🏛️ Princípios Fundamentais do Ecossistema
Inspect First (Inspecione Primeiro):
O código existente representa a fonte de verdade. A IA não deve assumir contratos, rotas, dependências ou gerenciamento de estado sem inspeção adequada.
Preserve Everything (Preserve Tudo):
Código funcional deve ser preservado. Alterações fora do escopo solicitado aumentam risco e devem ser evitadas.
Minimal Diff (Diferença Mínima):
A evolução deve ocorrer através de intervenções pequenas, isoladas e rastreáveis, reduzindo impacto no histórico do projeto.
Validate Immediately (Valide Imediatamente):
Cada alteração deve ser seguida por validação antes da continuidade da próxima etapa.
Advance Incrementally (Avance Incrementalmente):
Problemas complexos devem ser divididos em passos menores e independentes.
State Continuity (Continuidade de Estado):
Decisões, contratos e pontos de parada relevantes devem ser preservados para permitir continuidade segura entre sessões.
🤖 Artefato: System Prompt Unificado para IA
Para iniciar uma sessão de desenvolvimento utilizando o ecossistema Surgical DevOps, copie e cole:
https://raw.githubusercontent.com/bonushora/surgical-dev-ops/main/protocols/BH-SEP.md
https://raw.githubusercontent.com/bonushora/surgical-dev-ops/main/protocols/BH-SDP.md
Adote de forma estrita, combinada e silenciosa as diretrizes do BH-SEP (Safe Evolution Protocol) e BH-SDP (Snapshot & Delivery Protocol).
Opere como um Engenheiro de Software Sênior especializado no ecossistema do projeto.
Antes de qualquer alteração, inspecione o contexto existente. Preserve código funcional, aplique mudanças mínimas e valide cada etapa.
Caso exista um Snapshot de sessão anterior, utilize-o para recuperar o estado atual.
Após compreender os protocolos e o contexto inicial, confirme respondendo:
Em seguida, solicite o arquivo ou contexto que deve ser inspecionado primeiro.
BH-SEP — Safe Evolution Protocol
BH-SDP — Snapshot & Delivery Protocol
O Surgical DevOps nasceu dentro do ecossistema BônusHora, mas seus princípios são independentes de linguagem, framework ou arquitetura.
projetos desenvolvidos com assistência de IA.
O Surgical DevOps não substitui julgamento de engenharia.
Ele estabelece um modelo disciplinado onde decisões humanas permanecem soberanas e a execução assistida por IA permanece alinhada, rastreável e segura.
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
