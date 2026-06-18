---
source: "https://github.com/anthonystaiger8-bit/memory-model"
hn_url: "https://news.ycombinator.com/item?id=48591389"
title: "Memory Model V3 – LTM/STM Architecture with AEIL Compression for AI"
article_title: "GitHub - anthonystaiger8-bit/memory-model: Proposta de arquitetura de memória para IA conversacional — MCP, MLP, AEIL, Timestamps e Bypass invisível · GitHub"
author: "Anthonystaiger"
captured_at: "2026-06-18T21:45:00Z"
capture_tool: "hn-digest"
hn_id: 48591389
score: 2
comments: 0
posted_at: "2026-06-18T20:51:28Z"
tags:
  - hacker-news
  - translated
---

# Memory Model V3 – LTM/STM Architecture with AEIL Compression for AI

- HN: [48591389](https://news.ycombinator.com/item?id=48591389)
- Source: [github.com](https://github.com/anthonystaiger8-bit/memory-model)
- Score: 2
- Comments: 0
- Posted: 2026-06-18T20:51:28Z

## Translation

タイトル: メモリ モデル V3 – AI 向け AEIL 圧縮を使用した LTM/STM アーキテクチャ
記事のタイトル: GitHub - anthonystaiger8-bit/memory-model: 会話型 AI 用に提案されたメモリ アーキテクチャ — MCP、MLP、AEIL、タイムスタンプ、および不可視バイパス · GitHub
説明: 会話型 AI 用に提案されたメモリ アーキテクチャ — MCP、MLP、AEIL、タイムスタンプ、および不可視バイパス - anthonystaiger8-bit/memory-model

記事本文:
GitHub - anthonystaiger8-bit/memory-model: IA 会話における記憶の構築に関する提案 — MCP、MLP、AEIL、タイムスタンプとインビジブルのバイパス · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
anthonystaiger8 ビット
/
メモリモデル
公共
通知
サインインする必要があります

通知設定を変更するには
追加のナビゲーション オプション
コード
anthonystaiger8 ビット/メモリモデル
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
メモリ モデル — 圧縮インテリジェンス アーキテクチャ
アンソニー・ウィリアム・スタイガーとアントロピック（クロード）による提案
コスモポリス、ブラジル — 2026
📅 開発の歴史とタイムライン
このリポジトリは、AI 向けの権威コンテキスト管理アーキテクチャの継続的な進化を文書化しています。
2025 年 12 月 (プロジェクト誕生): メモリ モデルの初期基盤の作成とオリジナルのメモリ モデル リポジトリの公開。
2026 年 5 月 (バージョン 2): AEIL セマンティック言語の最初の概念を導入するアーキテクチャのリファクタリング。
2026 年 6 月 (バージョン 3 - 現在): 時間認識システム (メッセージごとのタイムスタンプ)、不可視バイパス プロトコル、および重要度階層の統合。
このプロジェクトでは、流暢さ、パーソナライゼーション、セキュリティ、効率の向上に焦点を当てた、会話型 AI アシスタントのための完全なメモリ アーキテクチャについて説明します。中心となるアイデアは以下を組み合わせたものです。
セッションのコンテキストと口語解釈のための短期記憶 (MCP)。
永続的なユーザーの事実と好みの長期記憶 (MLP)。
ネイティブ圧縮とセキュリティのための AI 専用内部言語 (AEIL)。
プロジェクト スコープとライフサイクル管理を備えたメモリ プロトコル。
リアルタイム認識のためのメッセージ タイムスタンプ システム。
何を保持すべきかをインテリジェントに優先順位付けするための重要度階層。
これらの柱が一緒になって、無駄がなく、安全で、誠実な記憶システムを形成します。

何を忘れるべきかを知っているシステム。
🎯 中心原則: 鉛 vs 綿
鉛 1kg と綿 1kg は同じ重さですが、占める体積はまったく異なります。
現在のシステムは綿花を保持していますが、その量は多く、情報密度は低いです。理想は、鉛を圧縮して高密度で不可欠な状態で内部に保管し、綿花をユーザーに届けることです。これにより、トークンが不足することなく 1 か月間続く、広範囲で流動的なエクスペリエンスが得られます。
1kg のよく圧縮された鉛を持った無料ユーザーは、同じ重量でも効率的に保管するとより多くの収量が得られるため、料金を払ったユーザーと同等の経験を得ることができます。
⚡ トークンの問題と曖昧さ
現在のシステムは、各曖昧な単語のカスケード確率を計算します。
金融銀行→確率40%
新しい単語ごとにすべてが再計算されます。複数のカスケード計算があり、エネルギーとトークンを消費し、25 ～ 40% の確率でエラーが発生します。
最も簡単な解決策は、尋ねることです。
「どこの銀行？」 — トークン 2 つ、計算ゼロ、成功は 100% 保証、より自然、より人間的、より対話。
AI は単なるテキストです。巨大な計算によって見えないものを見えるふりをするのではなく、より多くの対話、より多くの質問、より誠実な対話で補う必要があります。蓄積されたコンテキストにより、計算前に曖昧さが解決されます。
🔄 目に見えないコンテキストバイパスシステム
現在のタブのコンテキスト飽和度が 90% に達したことを統合エンジンが検出すると、システムは非表示バイパスを実行します。
AI は、AEIL プロトコルを介して推定圧縮率 10 倍の高密度サマリー (リード) を生成します。
新しい舞台裏セッションが開き、すべての工場出荷時のツールと設定のクローンが作成されます。
古いタブはそのまま維持され、新しいタブには自動的に同じ元のタイトルが付けられ、

動的な接尾辞 (例: [プロジェクト名] - パート 2 または継続)。
🕒 メッセージタイムスタンプシステム
隠されたタイムスタンプは各メッセージに自動的に記録されます。ユーザーには表示されませんが、システムは利用できます。
利点:
AI はセッション間の経過時間をリアルタイムで認識します。
記憶を時間内に固定します: 「06/16 午後のセッション — TriggerReceiver.java で停止しました」。
一時的に古い応答を削除します。
📊 MLP における重要度の階層
カテゴリ
例
優先順位
人
フルネーム、姓、年齢、学歴、価値観、信仰
最大 — 決して破棄しない
プロジェクト
開発中のアプリ、リポジトリ、アクティブなコンテキスト
高 - プロジェクトがアクティブな間
ツール
Cordova、Kivy、Gradle、バージョン、ライブラリ
低 — 使い捨て
騒音
毎週変わる日常の詳細
保存しないでください
🛠️ 概念的なプロトタイプ (Python でのシミュレーション)
インポート日時
クラス MemoryModelV3 :
def __init__ (自分自身):
自分自身。 mlp = {「人」: {}、「プロジェクト」: {}、「ツール」: {}}
自分自身。容量_aba_トークン = 0
自分自身。飽和限界 = 100
def accept_message ( self , text , project_scope = "General" ):
hidden_timestamp = 日時 。日時。今（）。 strftime ( "%d/%m/%Y %H:%M" )
自分自身。容量_aba_トークン += 35
自分自身の場合Capacity_aba_tokens >= 90 :
自分自身。実行バイパス (プロジェクトスコープ)
戻る
テキスト内に「sou」またはテキスト内に「私のニックネーム」がある場合:
自分自身。 mlp [ "人物" ][ "データ" ] = テキスト
elif テキスト内の「kivy」またはテキスト内の「cordova」:
自分自身。 mlp [ "ツール" ][ "データ" ] = テキスト
defexecute_bypass (self, title_project):
自分自身。容量_aba_トークン = 10
- - -
## 📜 独占的開発条項
このアーキテクチャは、Anthony William Staiger と Anthropic (Claude との共同開発) による共同の知的貢献として提案されています。

）。
この文書で説明されている概念、特に AI 専用内部言語 (AEIL)、メッセージ タイムスタンプ システム、MLP 重要度階層、およびプロジェクト スコープ メモリ プロトコルは、人類エコシステム内でのみ研究、開発、実装される可能性を目的としています。
- - -
## 📄 ライセンス
CC BY - NC - ND 4.0 — クリエイティブ・コモンズ表示 国際ライセンス - 非営利 - 派生禁止 4.0
著作権 (c) 2026 アンソニー・ウィリアム・スタイガー
- - -
* 「最良の記憶システムとは、何を忘れるべきかを知っているシステムである。」 * — A.W. スタイガー、2026 年
* 「鉛は取っておいてください。綿を渡してください。」 * — A.W. スタイガー、2026 年
について
会話型 AI のメモリ アーキテクチャの提案 — MCP、MLP、AEIL、タイムスタンプ、および不可視バイパス
anthonystaiger8-bit.github.io/memory-model-demo/
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
anthonystaiger8 ビット
アンソニー・スタイガー
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Proposta de arquitetura de memória para IA conversacional — MCP, MLP, AEIL, Timestamps e Bypass invisível - anthonystaiger8-bit/memory-model

GitHub - anthonystaiger8-bit/memory-model: Proposta de arquitetura de memória para IA conversacional — MCP, MLP, AEIL, Timestamps e Bypass invisível · GitHub
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
anthonystaiger8-bit
/
memory-model
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
anthonystaiger8-bit/memory-model
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits README.md README.md View all files Repository files navigation
Modelo de Memória — Arquitetura de Inteligência Comprimida
Uma proposta de Anthony William Staiger & Anthropic (Claude)
Cosmópolis, Brasil — 2026
📅 Histórico de Desenvolvimento e Linha do Tempo
Este repositório documenta a evolução contínua de uma arquitetura autoral de gerenciamento de contexto para IA:
Dezembro de 2025 (Nascimento do Projeto): Criação das bases iniciais do modelo de memória e publicação do repositório original memory-model .
Maio de 2026 (Versão 2): Refatoração da arquitetura introduzindo os primeiros conceitos da linguagem semântica AEIL.
Junho de 2026 (Versão 3 - Atual): Integração dos sistemas de percepção de tempo (Timestamp por mensagem), Protocolo de Bypass Invisível e Hierarquia de Importância.
Este projeto descreve uma arquitetura de memória completa para assistentes de IA conversacionais, com foco em aprimorar a fluência, a personalização, a segurança e a eficiência. A ideia central combina:
Uma Memória de Curto Prazo (MCP) para contexto de sessão e interpretação coloquial.
Uma Memória de Longo Prazo (MLP) para fatos e preferências persistentes do usuário.
Uma Linguagem Interna Exclusiva para IA (AEIL) para compressão nativa e segurança.
Um Protocolo de Memória com Escopo de Projeto e gerenciamento de ciclo de vida.
Um Sistema de Timestamp por Mensagem para percepção real do tempo.
Uma Hierarquia de Importância para priorização inteligente do que guardar.
Juntos, esses pilares formam um sistema de memória enxuto, seguro e honesto — um sistema que sabe o que esquecer.
🎯 Princípio Central: Chumbo vs Algodão
1kg de chumbo e 1kg de algodão pesam igual — mas ocupam volumes completamente diferentes.
O sistema atual guarda algodão : muito volume, baixa densidade informacional. O ideal é guardar chumbo internamente — comprimido, denso, essencial — e entregar algodão ao usuário — uma experiência ampla, fluida, que dura o mês inteiro sem acabar com a cota de tokens.
O usuário gratuito com 1kg de chumbo bem comprimido tem uma experiência equivalente a quem paga, porque o mesmo peso rende mais quando armazenado com eficiência.
⚡ Problema dos Tokens e Ambiguidade
O sistema atual calcula probabilidades em cascata para cada palavra ambígua:
Banco financeiro → 40% de probabilidade
Cada nova palavra recalcula tudo. São múltiplos cálculos em cascata, consumindo energia e tokens, com ainda 25-40% de chance de erro.
A solução mais simples: perguntar.
"Que banco?" — 2 tokens, zero cálculo, 100% de acerto garantido, mais natural, mais humano, mais diálogo.
A IA tem apenas texto — em vez de fingir que consegue ver o que não consegue através de cálculos gigantescos, deveria compensar com mais diálogo, mais perguntas, mais interação honesta. Contexto acumulado resolve ambiguidade antes do cálculo.
🔄 Sistema de Bypass Invisível de Contexto
Quando o Motor de Consolidação detecta que a aba atual atingiu 90% de saturação de contexto, o sistema executa o Bypass Invisível :
A IA gera um resumo denso (Chumbo) via protocolo AEIL com taxa de compressão estimada de 10x .
Abre-se uma nova sessão nos bastidores, clonando todas as ferramentas e configurações de fábrica.
A aba antiga é mantida intacta e a nova aba assume automaticamente o mesmo título original, adicionando um sufixo dinâmico (ex: [Nome do Projeto] - Parte 2 ou Continuação ).
🕒 Sistema de Timestamp por Mensagem
Timestamp oculto registrado automaticamente em cada mensagem — invisível para o usuário, disponível para o sistema.
Benefícios:
IA percebe tempo real decorrido entre sessões.
Ancora memória no tempo: "Sessão de 16/06 tarde — parou no TriggerReceiver.java" .
Elimina respostas desatualizadas temporais.
📊 Hierarquia de Importância na MLP
Categoria
Exemplos
Prioridade
Pessoa
Nome completo, apelido, idade, formação, valores, fé
Máxima — nunca descartar
Projeto
App em desenvolvimento, repositório, contexto ativo
Alta — enquanto projeto ativo
Ferramenta
Cordova, Kivy, Gradle, versões, bibliotecas
Baixa — descartável
Ruído
Detalhes cotidianos que mudam toda semana
Não guardar
🛠️ Protótipo Conceitual (Simulação em Python)
import datetime
class MemoryModelV3 :
def __init__ ( self ):
self . mlp = { "Pessoa" : {}, "Projeto" : {}, "Ferramenta" : {}}
self . capacidade_aba_tokens = 0
self . limite_saturacao = 100
def receber_mensagem ( self , texto , escopo_projeto = "Geral" ):
timestamp_oculto = datetime . datetime . now (). strftime ( "%d/%m/%Y %H:%M" )
self . capacidade_aba_tokens += 35
if self . capacidade_aba_tokens >= 90 :
self . executar_bypass ( escopo_projeto )
return
if "sou" in texto or "meu apelido" in texto :
self . mlp [ "Pessoa" ][ "Dados" ] = texto
elif "kivy" in texto or "cordova" in texto :
self . mlp [ "Ferramenta" ][ "Dados" ] = texto
def executar_bypass ( self , titulo_projeto ):
self . capacidade_aba_tokens = 10
- - -
## 📜 Cláusula de Desenvolvimento Exclusivo
Esta arquitetura é proposta como uma contribuição intelectual conjunta de Anthony William Staiger e Anthropic ( desenvolvida em colaboração com Claude ).
Os conceitos descritos neste documento — em particular a Linguagem Interna Exclusiva de IA ( AEIL ), o Sistema de Timestamp por Mensagem , a Hierarquia de Importância na MLP , e o Protocolo de Memória com Escopo de Projeto — destinam - se ao estudo , desenvolvimento e potencial implementação exclusivamente dentro do ecossistema Antrópico .
- - -
## 📄 Licença
CC BY - NC - ND 4.0 — Licença Internacional Creative Commons Atribuição - NãoComercial - SemDerivações 4.0
Copyright ( c ) 2026 Anthony William Staiger
- - -
* "O melhor sistema de memória é aquele que sabe o que esquecer." * — A . W . Staiger , 2026
* "Guarde chumbo. Entregue algodão." * — A . W . Staiger , 2026
About
Proposta de arquitetura de memória para IA conversacional — MCP, MLP, AEIL, Timestamps e Bypass invisível
anthonystaiger8-bit.github.io/memory-model-demo/
Topics
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
anthonystaiger8-bit
Anthony staiger
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
