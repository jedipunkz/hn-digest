---
source: "https://docs.cloud.google.com/gemini-enterprise-agent-platform/models/tuning/distillation?hl=de"
hn_url: "https://news.ycombinator.com/item?id=49086338"
title: "Gemini Distillation Service"
article_title: "Gemini Distillation Service | Gemini Enterprise Agent Platform | Google Cloud Documentation"
author: "theanonymousone"
captured_at: "2026-07-28T17:19:28Z"
capture_tool: "hn-digest"
hn_id: 49086338
score: 2
comments: 0
posted_at: "2026-07-28T16:27:55Z"
tags:
  - hacker-news
  - translated
---

# Gemini Distillation Service

- HN: [49086338](https://news.ycombinator.com/item?id=49086338)
- Source: [docs.cloud.google.com](https://docs.cloud.google.com/gemini-enterprise-agent-platform/models/tuning/distillation?hl=de)
- Score: 2
- Comments: 0
- Posted: 2026-07-28T16:27:55Z

## Translation

タイトル: ジェミニ蒸留サービス
記事のタイトル: ジェミニ蒸留サービス | Gemini エンタープライズ エージェント プラットフォーム | Google Cloud ドキュメント
説明: Gemini Distillation を使用して、より大きな教師モデルの出力と推論パスを使用して、より小規模で効率的な学生モデルをトレーニングする方法を学びます。

記事本文:
ジェミニ蒸留サービス | Gemini エンタープライズ エージェント プラットフォーム | Google Cloud ドキュメント
メインコンテンツにスキップ
技術分野
閉じる
AI と ML
分散型、ハイブリッド、マルチクラウド
可観測性と監視
アクセスとリソースの管理
コストと使用量の管理
SDK、言語、フレームワーク、ツール
Gemini エンタープライズ エージェント プラットフォーム
クロスプロダクトツール
もっと見る
エージェント プラットフォームの使用を開始する
Gen AI SDK を使用して Gemini API コードを開発する
アプリケーションのデフォルトの資格情報を構成する
Google AI Studio からエージェント プラットフォームへの移行
Gemini 3 プロンプト ガイド
OpenAI ライブラリを使用して Gemini モデルにアクセスする
モデルを選択してください
モデルガーデン
Model Garden でのモデルの使用
ジェミニ
最新の Gemini モデルに移行する
人間クロード モデルの割り当て
モデルガーデンよりパートナーモデルを提供
パートナーモデルの設定
埋め込み (e5)
多言語 E5 小
Googleジェマ
サービスとしてのモデル (MaaS)
手順: Gemma をデプロイし、推論を実行します (GPU)
手順: Gemma をデプロイし、推論を実行します (TPU)
マネージドオープンモデル（MaaS）
概要
Model as a Service (MaaS) 経由でオープン モデルを使用します。
開いているモデルへのアクセスを許可する
オープンモデルのMaaS APIを呼び出す
自己展開型オープンモデル
概要
オープンモデルを提供する
モデルガーデンからオープンモデルを提供
事前定義されたコンテナを使用してオープン モデルをデプロイする
カスタム vLLM コンテナを使用してオープン モデルをデプロイする
カスタムの重みを使用してモデルをデプロイする
ハグフェイスモデルを使用する
指示
Model Garden の高度な機能を使用してモデルのパフォーマンスを最適化
テキストおよびマルチモーダル LLM を配信するための vLLM の包括的なガイド (

GPU)
スポット VM と予約を使用して Llama 3 モデルを展開する
明確かつ具体的なガイダンスを提供する
コンテキスト情報を追加する
モデルにそのロジックを説明するように指示します。
複雑なタスクを分解する
パラメータ値を試してみる
迅速な反復のための戦略
プロンプトのタスク固有のガイダンス
マルチモーダル プロンプトを設計する
AI の責任ある使用
システムの安全上の注意事項
Gemini によるセキュリティ フィルタリングとコンテンツ モデレーション
ブロックされた返信を処理する
テキストとコードの生成
テキストの生成
コンテンツ生成パラメータ
画像生成
Gemini で画像を生成する
Gemini を使用してビデオから画像を生成する
Gemini を使用した画像生成のベスト プラクティス
Gemini での画像生成の制限
AI の責任ある適用と Gemini 画像生成の使用
最初のフレームからビデオを作成する
最初と最後のフレームからのビデオ
画像参照を含む画像要素で構成されるビデオ
Veo プロンプト リライタを無効にする
Veo 向け AI の責任ある適用
音楽の生成
ルリアの紹介
メディア分析
画像を理解する
境界ボックスの検出
Google 検索による基盤
検索 API を使用した基盤
並列 Web 検索を備えた基盤
Exa Web 検索による資金調達
行きましょう
Gen AI SDK の使用を開始する
ライブセッションの開始と管理
オーディオとビデオのストリームを送信する
言語と音声を設定する
Gemini 機能を構成する
ライブ API のベスト プラクティス
ライブAPIのバグ修正
テキストの埋め込み
テキスト埋め込みの取得
埋め込み用のタスク タイプを選択してください
マルチモーダル埋め込みの取得
バッチ埋め込み推論を取得する
AI を活用したツールを使用してプロンプトを作成する
概要
微調整の概要
G

emini モデルに適合
監視付き微調整
監視付き微調整
監視付き微調整を使用する
サポートされているモダリティ
テキストの最適化
強化学習に投票する
強化学習に投票する
強化学習の微調整を行う仕事
概要
優先投票
優先投票
継続的な最適化を使用する
LoRA および QLoRA との連携に関する推奨事項
オープンモデル
監視付きチューニングと蒸留チューニング
埋め込みモデル
テキスト埋め込みモデルを調整する
翻訳モデル
監視付き微調整
監視付き微調整を使用する
OpenAI ライブラリを使用したコール エージェント プラットフォーム モデル
概要
チュートリアル: コンソールを使用して評価を実行する
エージェント プラットフォーム SDK の GenAI クライアントを使用して評価を実行する
手順: エージェント プラットフォーム SDK の GenAI クライアントを使用してモデルを評価する
評価指標を定義する
評価指標を定義する
管理されたカテゴリベースのメトリクスの詳細
評価データセットを準備します。
評価結果の表示と解釈
代替の評価方法
Agent Platform SDKの評価モジュールを使用した評価
チュートリアル: Agent Platform SDK の評価モジュールを使用して評価を実行する
評価指標を定義する
評価データセットを準備します。
評価結果の解釈
モデルベースのメトリクスのテンプレート
計算ベースのスコアリング パイプラインを実行する
提供する
使用オプション – 概要
プロビジョニングされたスループット
プロビジョニングされたスループットの概要
プロビジョニングされたスループット要件の計算
ライブ API のプロビジョニングされたスループット
Gemini 3 モデルのプロビジョニングされたスループット (Nano Banana)
Veo 3 モデルのプロビジョニングされたスループット
準備完了のジェスチャー

単一ゾーンの l 番目のスループット
プロビジョニングされたスループットを使用する
クラウドストレージからバッチジョブを作成する
BigQuery からバッチジョブを作成する
未完了のバッチ ジョブを再開する
キャッシュ内のプロンプト コンテキストを再利用する
概要
コンテキストキャッシュ情報の取得
微調整された Gemini モデルのコンテキスト キャッシュ
生成AIモデルを導入する
Model Garden モデルへのアクセスを制御する
データアクセスの監査ログを有効にする
カスタムメタデータラベルを使用してコストを追跡する
IAP を使用した安全な生成 AI アプリ
概要
プロジェクトとソースリポジトリをセットアップする
独自のモデルを作成する
概要
エージェントプラットフォームのインターフェース
エージェント プラットフォームの初心者ガイド
AutoML モデルをトレーニングする
カスタム モデルをトレーニングする
カスタムモデルから推論を取得する
エージェント プラットフォームと Python SDK を使用してモデルをトレーニングする
はじめに
統合された ML フレームワーク
パイトーチ
BigQuery ユーザー向けのエージェント プラットフォーム
プロジェクトと開発環境をセットアップする
Python 用エージェント プラットフォーム SDK をインストールする
エージェント プラットフォームに対する認証
指示を読む
指示の概要
AutoML ガイド
こんにちは画像データ
概要
プロジェクトと環境をセットアップする
データセットを作成して画像をインポートする
AutoML 画像分類モデルをトレーニングする
モデルのパフォーマンスを評価および分析する
モデルをエンドポイントにデプロイして推論を実行する
表形式の Hello データ
概要
プロジェクトと環境をセットアップする
データセットを作成して AutoML 分類モデルをトレーニングする
モデルをデプロイして推論をリクエストする
カスタムトレーニングガイド
カスタム表形式モデルをトレーニングする
TensorFlow Keras 画像分類モデルをトレーニングする
概要
プロジェクトと環境をセットアップする
いる

カスタム画像分類モデルをトレーニングする
カスタム画像分類モデルからの予測を提供する
カスタムデータを使用して画像分類モデルを最適化する
エージェント プラットフォーム開発ツールを使用する
エージェント プラットフォーム SDK を使用する
概要
Python 用エージェント プラットフォーム SDK の紹介
Python クラス用のエージェント プラットフォーム SDK
エージェント プラットフォーム SDK クラスの概要
エージェント プラットフォームの Terraform サポート
エージェント プラットフォームでのサーバーレス トレーニング
エージェント プラットフォームでのサーバーレス トレーニングの概要
データをロードして準備する
データ準備の概要
Cloud Storage をマウントされたファイル システムとして使用する
サーバーレス トレーニング用の NFS 共有の展開
トレーニング申請書を準備する
サーバーレス研修サービスについて
事前定義されたコンテナを使用する
事前定義されたコンテナ用の Python トレーニング アプリケーションを作成する
サーバーレストレーニング用の事前定義されたコンテナー
カスタムコンテナを使用する
サーバーレストレーニング用のカスタムコンテナ
カスタムコンテナの作成
トレーニング コードをローカルでコンテナ化して実行する
不揮発性リソースでトレーニングする
概要
永続的なリソースを作成する
永続リソース上でトレーニング ジョブを実行する
不揮発性リソースに関する情報の取得
永続リソースを再起動します
不揮発性リソースの削除
トレーニング ジョブを構成する
カスタムトレーニング方法を選択してください
トレーニング用のコンテナ設定を構成する
トレーニング用のコンピューティング リソースを構成する
トレーニングで予約を使用する
トレーニングにスポット VM を使用する
トレーニング ジョブを送信する
カスタムジョブの作成
ハイパーパラメータ調整
ハイパーパラメータファイン

投票する
ハイパーパラメータ微調整を使用する
リソースの可用性に基づいてジョブをスケジュールする
カスタム トレーニングにプライベート IP アドレスを使用する
トレーニングには Private Service Connect インターフェイスを使用する (推奨)
監視とトラブルシューティング
インタラクティブシェルを使用したトレーニングの監視とトラブルシューティング
プロファイルモデルのトレーニングパフォーマンス
方法: 継続的なトレーニング用のパイプラインを作成する
組織ポリシーのカスタム制限を作成する
Vertex AI トレーニング クラスター
概要
トレーニングクラスターの開始
導入に関する考慮事項
リソースを計算する
クラスターの作成と管理
クラスターの作成
クラスター内のアカウントとジョブのスケジュールを管理する
機能ガイド
Slurm クラスターで Flex Start VM を使用する
クラスター内でワークロードを実行する
事前に構築されたワークロードを実行する
TensorBoard を使用してジョブを視覚化する
エージェントプラットフォームのレイ
Ray on Agent プラットフォームの概要
エージェント プラットフォームでの Ray のセットアップ
エージェント プラットフォームで Ray クラスターを作成する
エージェント プラットフォーム上の Ray クラスターを監視する
エージェント プラットフォームでの Ray クラスターのスケーリング
エージェント プラットフォームで Ray アプリケーションを開発する
エージェント プラットフォームで Spark on Ray クラスターを実行する
BigQuery を使用したエージェント プラットフォームでの Ray の使用
モデルをデプロイして推論を取得する
ニューラル アーキテクチャ検索の実行
概要
PyTorch のトレーニング速度を最適化する
事前定義されたトレーニング コンテナと検索領域を使用する
Agent Platform Vizier による最適化
エージェント プラットフォーム Vizier の概要
エージェント プラットフォーム Vizier スタディの作成
AutoML モデルの開発
AutoML トレーニングの概要
画像データ
分類
データの準備
物体検出
データの準備
AutoML Edge モデルのエクスポート

表形式データの概要
表形式のワークフロー
概要
AutoML を使用して分類と回帰を実行する
概要
クイック スタート ガイド: AutoML 分類（Cloud コンソール）
AutoML を使用して予測を作成する
概要
ARIMA+ で予測を作成する
Prophet で予測を実行する
分類と回帰のための特徴の属性
予測のための特徴の属性
AutoML 表形式データのデータ型とデータ変換
予測用のトレーニングパラメータ
テーブルデータのデータ分割
表形式のトレーニング データを作成するためのベスト プラクティス
AutoML Edge モデルをトレーニングする
コンソールを使用する
生成AIモデルの開発
記録の作成と管理
AutoML モデルのデータ分割
データセットからメタデータとアノテーションをエクスポートする
画像データセットのバージョンを管理する (API のみ)
推論用のモデルを構成する
推論用にモデル アーティファクトをエクスポートする
推論用の事前定義されたコンテナ
推論のためのカスタムコンテナ要件
推論にカスタム コンテナを使用する
カスタム ルートを使用する
最適化された TensorFlow ランタイムを使用する
NVIDIA Triton による推論の提供
カスタム推論ルーチン
オンライン推論の取得
エンドポイントの作成
エンドポイントの種類を選択してください
パブリックエンドポイントer

[切り捨てられた]

## Original Extract

Hier erfahren Sie, wie Sie mit Gemini Distillation kleinere, effiziente Schülermodelle mithilfe der Ausgabe und der Begründungspfade größerer Lehrermodelle trainieren können.

Gemini Distillation Service | Gemini Enterprise Agent Platform | Google Cloud Documentation
Zum Hauptinhalt springen
Technologiebereiche
close
KI und ML
Verteilt, Hybrid und Multi-Cloud
Beobachtbarkeit und Monitoring
Zugriffs- und Ressourcenverwaltung
Kosten- und Nutzungsmanagement
SDK, Sprachen, Frameworks und Tools
Gemini Enterprise Agent Platform
Produktübergreifende Tools
Mehr
Erste Schritte mit der Agent Platform
Gemini API-Code mit dem Gen AI SDK entwickeln
Standardanmeldedaten für Anwendungen konfigurieren
Von Google AI Studio zur Agent Platform migrieren
Leitfaden für Gemini 3-Prompts
Mit OpenAI-Bibliotheken auf Gemini-Modelle zugreifen
Modelle auswählen
Model Garden
Modelle in Model Garden verwenden
Gemini
Auf die neuesten Gemini-Modelle migrieren
Kontingente für Anthropic Claude-Modelle
Partnermodelle aus Model Garden bereitstellen
Einstellung von Partnermodellen
Einbettung (e5)
Mehrsprachig E5 Small
Google Gemma
Model-as-a-Service (MaaS)
Anleitung: Gemma bereitstellen und Inferenz durchführen (GPU)
Anleitung: Gemma bereitstellen und Inferenz durchführen (TPU)
Verwaltete offene Modelle (MaaS)
Übersicht
Offene Modelle über Model as a Service (MaaS) verwenden
Zugriff auf offene Modelle gewähren
MaaS-APIs für offene Modelle aufrufen
Selbst bereitgestellte offene Modelle
Übersicht
Offene Modelle bereitstellen
Offene Modelle aus Model Garden bereitstellen
Offene Modelle mit vordefinierten Containern bereitstellen
Offene Modelle mit einem benutzerdefinierten vLLM-Container bereitstellen
Modelle mit benutzerdefinierten Gewichtungen bereitstellen
Hugging Face-Modelle verwenden
Anleitungen
Modellleistung mit erweiterten Funktionen in Model Garden optimieren
Umfassender Leitfaden zu vLLM für die Bereitstellung von Text- und multimodalen LLMs (GPU)
Llama 3-Modelle mit Spot-VMs und Reservierungen bereitstellen
Klare und spezifische Anleitung geben
Kontextinformationen hinzufügen
Weisen Sie das Modell an, seine Logik zu erklären.
Komplexe Aufgaben aufschlüsseln
Mit Parameterwerten experimentieren
Strategien für Prompt-Iteration
Aufgabenspezifische Anleitung für Prompts
Multimodale Prompts entwerfen
Verantwortungsbewusste Anwendung von KI
Systemanweisungen zur Sicherheit
Gemini für Sicherheitsfilterung und Inhaltsmoderation
Blockierte Antworten verarbeiten
Text- und Codegenerierung
Textgenerierung
Parameter für die Inhaltsgenerierung
Bildgenerierung
Bilder mit Gemini generieren
Bilder aus Videos mit Gemini generieren
Best Practices für die Bildgenerierung mit Gemini
Einschränkungen bei der Bildgenerierung mit Gemini
Verantwortungsbewusste Anwendung von KI und Nutzung der Gemini-Bildgenerierung
Video aus dem ersten Frame erstellen
Video aus erstem und letztem Frame
Video aus Bildelementen mit Bildreferenzen
Prompt-Rewriter von Veo deaktivieren
Verantwortungsbewusste Anwendung von KI für Veo
Musikgenerierung
Einführung in Lyria
Media-Analyse
Bilder verstehen
Erkennung von Begrenzungsrahmen
Fundierung mit der Google Suche
Fundierung mit Ihrer Search API
Fundierung mit Parallel Web Search
Fundierung mit der Exa-Websuche
Los gehts
Erste Schritte mit dem Gen AI SDK
Live-Sitzungen starten und verwalten
Audio- und Videostreams senden
Sprache und Stimme konfigurieren
Gemini-Funktionen konfigurieren
Best Practices für die Live API
Fehlerbehebung bei der Live API
Texteinbettungen
Texteinbettungen abrufen
Aufgabentyp für Einbettungen auswählen
Multimodale Einbettungen abrufen
Batch-Embeddings-Inferenz abrufen
KI-gestützte Tools zum Schreiben von Prompts verwenden
Übersicht
Einführung in die Feinabstimmung
Gemini-Modelle abstimmen
Überwachte Feinabstimmung
Überwachte Feinabstimmung
Überwachte Feinabstimmung verwenden
Unterstützte Modalitäten
Textoptimierung
Abstimmung für Reinforcement Learning
Abstimmung für Reinforcement Learning
Job für die Feinabstimmung für Reinforcement Learning
Übersicht
Präferenzabstimmung
Präferenzabstimmung
Kontinuierliche Optimierung verwenden
Empfehlungen für die Abstimmung mit LoRA und QLoRA
Offene Modelle
Überwachte Feinabstimmung und Destillations-Feinabstimmung
Einbettungsmodelle
Texteinbettungsmodelle abstimmen
Übersetzungsmodelle
Überwachte Feinabstimmung
Überwachte Feinabstimmung verwenden
Agent Platform-Modelle mit OpenAI-Bibliotheken aufrufen
Übersicht
Tutorial: Bewertung mit der Console durchführen
Bewertung mit dem GenAI-Client im Agent Platform SDK durchführen
Anleitung: Modelle mit dem GenAI-Client im Agent Platform SDK bewerten
Bewertungsmesswerte definieren
Bewertungsmesswerte definieren
Details zu verwalteten rubrikbasierten Messwerten
Bereiten Sie das Bewertungs-Dataset vor.
Bewertungsergebnisse ansehen und interpretieren
Alternative Bewertungsmethoden
Bewertung mit dem Bewertungsmodul im Agent Platform SDK
Tutorial: Bewertung mit dem Bewertungsmodul im Agent Platform SDK durchführen
Bewertungsmesswerte definieren
Bereiten Sie das Bewertungs-Dataset vor.
Bewertungsergebnisse interpretieren
Vorlagen für modellbasierte Messwerte
Berechnungsbasierte Bewertungspipeline ausführen
Bereitstellen
Nutzungsoptionen – Übersicht
Bereitgestellter Durchsatz
Übersicht über den bereitgestellten Durchsatz
Anforderungen an Bereitgestellten Durchsatz berechnen
Bereitgestellter Durchsatz für Live API
Bereitgestellter Durchsatz für Gemini 3-Modelle (Nano Banana)
Bereitgestellter Durchsatz für Veo 3-Modelle
Bereitgestellter Durchsatz für eine einzelne Zone
Provisioned Throughput verwenden
Batchjob aus Cloud Storage erstellen
Batchjob aus BigQuery erstellen
Unvollständigen Batch-Job fortsetzen
Prompt-Kontext im Cache wiederverwenden
Übersicht
Informationen zum Kontext-Cache abrufen
Kontext-Cache für feinabgestimmte Gemini-Modelle
Generative KI-Modelle bereitstellen
Zugriff auf Model Garden-Modelle steuern
Audit-Logs zum Datenzugriff aktivieren
Kosten mithilfe von benutzerdefinierten Metadatenlabels im Blick behalten
Generative KI-App mit IAP sichern
Übersicht
Projekt und Quell-Repository einrichten
Eigenes Modell erstellen
Übersicht
Schnittstellen für die Agent Platform
Anleitungen für Einsteiger für die Agent Platform
AutoML-Modell trainieren
Benutzerdefiniertes Modell trainieren
Inferenz von einem benutzerdefinierten Modell abrufen
Modell mit der Agent Platform und dem Python SDK trainieren
Einführung
Integrierte ML-Frameworks
PyTorch
Agent Platform für BigQuery-Nutzer
Projekt- und Entwicklungsumgebung einrichten
Agent Platform SDK für Python installieren
Bei der Agent Platform authentifizieren
Anleitung lesen
Übersicht über Anleitungen
AutoML-Anleitungen
Hello-Bilddaten
Übersicht
Projekt und Umgebung einrichten
Dataset erstellen und Bilder importieren
AutoML-Bildklassifizierungsmodell trainieren
Modellleistung bewerten und analysieren
Modell auf einem Endpunkt bereitstellen und eine Inferenz durchführen
Tabellarische Hello-Daten
Übersicht
Projekt und Umgebung einrichten
Dataset erstellen und AutoML-Klassifizierungsmodell trainieren
Modell bereitstellen und Inferenz anfordern
Benutzerdefinierte Trainingsanleitungen
Benutzerdefiniertes tabellarisches Modell trainieren
TensorFlow Keras-Bildklassifizierungsmodell trainieren
Übersicht
Projekt und Umgebung einrichten
Benutzerdefiniertes Bildklassifizierungsmodell trainieren
Vorhersagen aus einem benutzerdefinierten Bildklassifizierungsmodell bereitstellen
Bildklassifizierungsmodell mit benutzerdefinierten Daten optimieren
Entwicklungstools der Agent Platform verwenden
Agent Platform SDK verwenden
Übersicht
Einführung in das Agent Platform SDK für Python
Agent Platform SDK für Python-Klassen
Agent Platform SDK-Klassen – Übersicht
Terraform-Support für die Agent Platform
Serverloses Training auf der Agent Platform
Übersicht über das serverlose Training in Agent Platform
Daten laden und vorbereiten
Übersicht über die Datenvorbereitung
Cloud Storage als bereitgestelltes Dateisystem verwenden
NFS-Freigabe für serverloses Training bereitstellen
Trainingsanwendung vorbereiten
Informationen zum serverlosen Trainingsdienst
Vordefinierte Container verwenden
Python-Trainingsanwendung für einen vordefinierten Container erstellen
Vordefinierte Container für serverloses Training
Benutzerdefinierte Container verwenden
Benutzerdefinierte Container für serverloses Training
Benutzerdefinierten Container erstellen
Trainingscode lokal containerisieren und ausführen
Auf einer nichtflüchtigen Ressource trainieren
Übersicht
Nichtflüchtige Ressource erstellen
Trainingsjobs für eine nichtflüchtige Ressource ausführen
Informationen zu nichtflüchtigen Ressourcen abrufen
Nichtflüchtige Ressource neu starten
Nichtflüchtige Ressource löschen
Trainingsjob konfigurieren
Benutzerdefinierte Trainingsmethode auswählen
Containereinstellungen für das Training konfigurieren
Rechenressourcen für das Training konfigurieren
Reservierungen mit Training verwenden
Spot-VMs für das Training verwenden
Trainingsjob senden
Benutzerdefinierte Jobs erstellen
Hyperparameter-Abstimmung
Hyperparameter-Feinabstimmung
Hyperparameter-Feinabstimmung verwenden
Jobs basierend auf der Ressourcenverfügbarkeit planen
Private IP-Adresse für benutzerdefiniertes Training verwenden
Private Service Connect-Schnittstelle für das Training verwenden (empfohlen)
Monitoring und Fehlerbehebung
Training mit interaktiver Shell überwachen und Fehler beheben
Leistung des Profilmodelltrainings
Anleitung: Pipeline für kontinuierliches Training erstellen
Benutzerdefinierte Einschränkungen für Organisationsrichtlinien erstellen
Vertex AI-Trainingscluster
Übersicht
Erste Schritte mit Trainingsclustern
Überlegungen zur Bereitstellung
Ressourcen berechnen
Cluster erstellen und verwalten
Cluster erstellen
Konten und Jobplanung in einem Cluster verwalten
Leitfäden für Funktionen
Flex-Start-VMs mit Slurm-Clustern verwenden
Arbeitslast in einem Cluster ausführen
Vorgefertigte Arbeitslasten ausführen
Jobs mit TensorBoard visualisieren
Ray in Agent Platform
Ray on Agent Platform – Übersicht
Einrichtung für Ray in Agent Platform
Ray-Cluster in der Agent Platform erstellen
Ray-Cluster auf der Agent Platform überwachen
Ray-Cluster in Agent Platform skalieren
Ray-Anwendung auf der Agent Platform entwickeln
Spark auf Ray-Cluster in Agent Platform ausführen
Ray in der Agent Platform mit BigQuery verwenden
Modell bereitstellen und Inferenzen abrufen
Neural Architecture Search durchführen
Übersicht
Trainingsgeschwindigkeit für PyTorch optimieren
Vordefinierte Trainingscontainer und Suchbereiche verwenden
Mit Agent Platform Vizier optimieren
Übersicht über Agent Platform Vizier
Agent Platform Vizier-Studien erstellen
Entwicklung von AutoML-Modellen
Übersicht über das AutoML-Training
Bilddaten
Klassifizierung
Daten vorbereiten
Objekterkennung
Daten vorbereiten
AutoML Edge-Modell exportieren
Einführung in tabellarische Daten
Tabellarische Workflows
Übersicht
Klassifizierung und Regression mit AutoML durchführen
Übersicht
Kurzanleitung: AutoML-Klassifizierung (Cloud Console)
Prognosen mit AutoML erstellen
Übersicht
Prognosen mit ARIMA+ erstellen
Prognosen mit Prophet ausführen
Feature-Attributionen für Klassifizierung und Regression
Feature-Attributionen für Prognosen
Datentypen und Datenumwandlungen für tabellarische AutoML-Daten
Trainingsparameter für die Prognose
Datenaufteilungen für Tabellendaten
Best Practices für das Erstellen von tabellarischen Trainingsdaten
AutoML Edge-Modell trainieren
Console verwenden
Entwicklung von generativen KI-Modellen
Datensätze erstellen und verwalten
Datenaufteilungen für AutoML-Modelle
Metadaten und Annotationen aus einem Dataset exportieren
Versionen von Bild-Datasets verwalten (nur API)
Modelle für die Inferenz konfigurieren
Modellartefakte für die Inferenz exportieren
Vordefinierte Container für die Inferenz
Anforderungen an benutzerdefinierte Container für die Inferenz
Benutzerdefinierten Container für die Inferenz verwenden
Beliebige benutzerdefinierte Routen verwenden
Optimierte TensorFlow-Laufzeit verwenden
Inferenz mit NVIDIA Triton bereitstellen
Benutzerdefinierte Inferenzroutinen
Onlineinferenzen abrufen
Endpunkt erstellen
Endpunkttyp auswählen
Öffentlichen Endpunkt er

[truncated]
