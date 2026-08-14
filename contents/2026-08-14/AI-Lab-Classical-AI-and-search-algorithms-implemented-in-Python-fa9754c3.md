---
source: "https://github.com/ob22a/ai-lab"
hn_url: "https://news.ycombinator.com/item?id=49304392"
title: "AI-Lab: Classical AI and search algorithms implemented in Python"
article_title: "GitHub - ob22a/ai-lab: A highly modular Object-Oriented AI framework and interactive lab implemented in Python. Features automated benchmarking, pluggable solver engines, and interactive Pygame visualizers for Search, CSPs, Optimization, and Adversarial Games. · GitHub"
author: "ob22a"
captured_at: "2026-08-14T21:16:36Z"
capture_tool: "hn-digest"
hn_id: 49304392
score: 1
comments: 0
posted_at: "2026-08-14T20:47:57Z"
tags:
  - hacker-news
  - translated
---

# AI-Lab: Classical AI and search algorithms implemented in Python

- HN: [49304392](https://news.ycombinator.com/item?id=49304392)
- Source: [github.com](https://github.com/ob22a/ai-lab)
- Score: 1
- Comments: 0
- Posted: 2026-08-14T20:47:57Z

## Translation

タイトル: AI-Lab: Python で実装された古典的な AI と検索アルゴリズム
記事のタイトル: GitHub - ob22a/ai-lab: Python で実装された高度にモジュール化されたオブジェクト指向 AI フレームワークとインタラクティブなラボ。自動ベンチマーク、プラグイン可能なソルバー エンジン、検索、CSP、最適化、敵対ゲーム向けのインタラクティブな Pygame ビジュアライザーを備えています。 · GitHub
説明: 高度にモジュール化されたオブジェクト指向 AI フレームワークと、Python で実装されたインタラクティブなラボ。自動ベンチマーク、プラグイン可能なソルバー エンジン、検索、CSP、最適化、敵対ゲーム向けのインタラクティブな Pygame ビジュアライザーを備えています。 - ob22a/ai-lab

記事本文:
GitHub - ob22a/ai-lab: 高度にモジュール化されたオブジェクト指向 AI フレームワークと、Python で実装されたインタラクティブなラボ。自動ベンチマーク、プラグイン可能なソルバー エンジン、検索、CSP、最適化、敵対ゲーム向けのインタラクティブな Pygame ビジュアライザーを備えています。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ob22a
/
アイラボ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
102 コミット 102 コミット ベンチマーク ベンチマーク co

re core csp csp デモ デモ ドメイン ドメイン ゲーム ゲーム pdbs pdbs レポート レポート 結果 結果 スクリーンショット スクリーンショット 検索 検索テスト テスト ユーティリティ ユーティリティ 視覚化 視覚化 .gitignore .gitignore ライセンス ライセンス README.md README.md main.py main.py 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイル ナビゲーション
🚀 古典的な AI アルゴリズム フレームワークとインタラクティブ ラボ
Python でゼロから実装された、最先端の高度にモジュール化されたオブジェクト指向 AI フレームワーク。視覚学習者、研究者、開発者向けに特別に作成されたこのリポジトリは、膨大な数の人工知能パラダイムにわたって、統合された抽象化、プラグイン可能なソルバー エンジン、100% 透明なビジュアライザー、自動ベンチマーク パイプライン、およびインタラクティブな CLI ランチャー ( main.py ) を提供します。
# 依存関係をインストールする
pip インストール pygame matplotlib numpy
# インタラクティブな CLI メニューを起動します
Python main.py
# またはデモを直接実行する
python -m デモ.maze --algo AStar --vis
python -m Demon.local_search_tsp --algo GeneticAlgorithm --vis
python -m デモ.crazy_demo
🎮 統合インタラクティブ ランチャー ハブ ( main.py )
インタラクティブな CLI メニューを起動して、サポートされている 18 個の AI デモをすべて参照して実行します。
Python main.py
ランチャーハブの特徴:
カテゴリ別メニュー: 検索、最適化、CSP、および敵対ゲームのカテゴリにわたるデモを参照します。
アルゴリズム セレクター : 各デモは、アルゴリズム (A*、BFS、DFS、UCS、IDA*、ヒル クライミング、シミュレーテッド アニーリング、遺伝的アルゴリズム、バックトラッキング + MRV/MAC、Minimax、AlphaBeta、MCTS、IS-MCTS) を選択するための CLI フラグをサポートしています。
ゲームモード：人間対AI、AI対AI、人間対人間。
ビジュアライザーのサポート : --vis フラグが付いているすべてのデモは、インタラクティブな Pygame ビジュアライザーを起動します。
動的なウィンドウのサイズ変更: すべてのビジュアライザーで pygame.RESIZABLE を完全にサポートします。
📸 インタラクティブビジュアライザー G

ギャラリーとスクリーンショットのショーケース
🧩 制約満足度およびグラフ分解
ツリー分解 (ジャンクション ツリーとセパレーター)
サイクルカットセット条件付け (非循環ツリー部分問題)
数独 CSP (バックトラッキング + MRV + MAC)
N-Queens CSP 対称性の破れ
🧭 グラフと経路探索検索
迷路A*探索(マンハッタン/ユークリッド)
オンライン迷路検索 (LRTA* リアルタイム学習)
8 パズル スライディング タイル検索 (素の PDB)
ルーマニア地図の都市ルート
倉庫番ボックス押し検索
倉庫番解決状態
📈 継続的な最適化と母集団ソルバー
TSP 遺伝的アルゴリズム (エリート集団検査官)
TSP シミュレーテッドアニーリング (距離進行曲線)
N-Queens 遺伝的アルゴリズム (上位染色体)
N-Queens ローカル ビーム検索 (k 個の平行ビーム)
N-Queens 模擬アニーリング
N-クイーンズ ヒル クライミング
🎮 敵対ゲーム理論とカードゲーム
クレイジー カード ゲーム (情報セット MCTS 対 Obssa のヒューリスティック)
オセロ / リバーシ (アルファ-ベータ枝刈り)
Connect Four (アルファ-ベータ検索)
チェッカー (ミニマックス & アルファベータ検索)
クレイジーカードゲーム（情報セットMCTS）
三目並べ (ミニマックス検索)
🛠️ 教育モジュール、スタンドアロン デモ、および CLI の実行
このラボは、実践的な学習、実験、研究のために構築されています。ターミナルから直接デモを実行します。
# 1. グラフと経路探索の検索
python -m デモ.maze --algo AStar --vis
python -m デモ.maze --algo IGBFS --vis
python -m Demon.n-puzzle --size 4 --algo AStar --vis
python -mデモ.romanian_map_demo --start Arad --goal Bucharest --algo AStar --vis
python -m デモ.sokoban_demo --vis
# 2. ローカル検索と継続的な最適化
python -m Demon.local_search_tsp --algo GeneticAlgorithm --vis
python -mデモ.local_search_tsp --algo LocalBeamSearch --vis
python -m デモ.local_search_nqueens --algo GeneticAlgorithm --vis
#3. 制約満足問題 (CSP)
パイソン

-m デモ.csp_tree_decomposition --vis
python -m デモ.csp_cycle_cutset --vis
python -m デモ.csp_sudoku --difficulty ハード --inference mac --vis
python -m デモ.csp_map_coloring --vis
python -m デモ.csp_cryptarithmetic --vis
#4. ボードゲーム＆不完全情報カードゲーム
python -m デモ.games_demo --ゲーム オセロ --p1 人間 --p2 アルファベット --vis
python -mデモ.games_demo --game connect_four --p1 mcts --p2 random --vis
python -m デモ.crazy_demo
🏗 アーキテクチャとデザインシステム
このラボの核となる設計哲学は、ドメイン (状態表現) とソルバー (アルゴリズム) の間の懸念事項の明確な分離を中心に展開しています。
クラス図
クラス Search問題 {
+開始
+目標
+get_actions(状態)
+get_result(状態、アクション)
+get_cost(状態、アクション、next_state)
+ヒューリスティック(状態)
}
クラス最適化問題 {
+初期状態
+値(状態)
+get_all_neighbors(状態)
+get_random_neighbor(状態)
+クロスオーバー(状態1、状態2)
+変異(状態)
}
クラス CSP問題 {
+変数
+ドメイン
+制約
+add_constraint(制約)
}
クラス GameState {
+現在のプレイヤー
+get_legal_actions()
+apply_action(アクション)
+is_terminal()
+get_utility(プレイヤー)
}
クラスMazeSearch問題
クラスNパズル問題
クラスルーマニア語Map問題
クラス WordLadder問題
クラス倉庫番問題
クラス VacuumWorld問題
クラス NQueens問題
クラスTSP問題
クラスMapColoringCSP
クラスNQueensCSP
クラス SudokuCSP
クラス CryptarithmeticCSP
クラスTimetablingCSP
クラスTicTacToeState
クラスConnectFourState
クラス CheckersState
クラスオセロステート
クラスCrazyState
検索問題 <|-- 迷路検索問題
検索問題 <|-- NPパズル問題
検索問題 <|-- ルーマニア語地図問題
Search問題 <|-- WordLadder 問題
Search問題 <|-- 倉庫番問題
Search問題 <|-- VacuumWorld問題
最適化問題 <|-- NQueens 問題
最適化問題 <|-- TSPP

問題
CSP問題 <|-- MapColoringCSP
CSP問題 <|-- NQueensCSP
CSP問題 <|-- SudokuCSP
CSP問題 <|-- CryptarithmeticCSP
CSP問題 <|-- 時刻表CSP
ゲームの状態 <|-- TicTacToe の状態
ゲーム状態 <|-- ConnectFourState
ゲーム状態 <|-- チェッカー状態
ゲーム状態 <|-- オセロ状態
ゲームステート <|-- クレイジーステート
読み込み中
📊 パフォーマンスのベンチマークとレポートの生成
ベンチマーク評価結果とパフォーマンス比較レポートは、マークダウン形式と高解像度のチャート形式で直接保存されます。
# すべてのベンチマークを実行します (アルゴリズムごとに 30 回の反復)
python -m benchmarks.run_all_benchmarks --runs 30
# フィルターを使用して個別のベンチマーク スイートを実行する
python -m benchmarks.search_benchmark --runs 10 --domains 8pzl --algos " A*,IDA* "
python -m benchmarks.csp_benchmark --runs 10 --algos " BT+MAC,BT+MRV "
python -m benchmarks.game_benchmark --runs 5 --games tic_tac_toe
python -m benchmarks.local_search_benchmark --runs 10 --domains tsp
# 書き込む前に --reset を使用して既存の CSV データをクリアします
python -m benchmarks.search_benchmark --runs 30 --reset
レポートの生成
# パフォーマンス チャートを生成 (reports/figures/ に保存)
python -m ベンチマーク.generate_report
# マークダウン概要テーブルを生成 (reports/benchmark_report.md を更新)
python -m utils.generate_markdown_tables
レポート : reports/benchmark_report.md および reports/comparison.md を参照してください。
CSV 結果: results/*.csv 内の実行ごとの生データ。
チャート : レポート/図/ の高解像度の図。
🤝 オープンソース貢献ガイドライン
オープンソースの貢献を歓迎します。新しい検索アルゴリズム、ヒューリスティック、またはドメインを追加するのは簡単です。
新しい検索アルゴリズムを追加します。 search/SearchAlgorithm.py の SearchAlgorithm から継承し、 search_step() を実装します。
新しい CSP ヒューリスティック / 推論の追加 : c 内に (csp、代入) を受け取る関数を実装します。

sp/heuristics/ または csp/inference/ 。
新しいゲーム ドメインを追加します。 games/GameState.py の GameState から継承し、 get_legal_actions() 、 apply_action() 、および is_terminal() を定義します。
カスタム ビジュアライザーの追加 : 標準化された HUD コントロール (SPACE 自動再生、+ / - 速度、LEFT / RIGHT ステップ、R 再起動) を使用して、視覚化/に Pygame ビジュアライザー クラスを作成します。
MITライセンスに基づいてライセンスされています。教育研究、視覚学習、高度な AI アルゴリズムの探索のために開発されました。
高度にモジュール化されたオブジェクト指向 AI フレームワークと、Python で実装されたインタラクティブなラボ。自動ベンチマーク、プラグイン可能なソルバー エンジン、検索、CSP、最適化、敵対ゲーム向けのインタラクティブな Pygame ビジュアライザーを備えています。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A highly modular Object-Oriented AI framework and interactive lab implemented in Python. Features automated benchmarking, pluggable solver engines, and interactive Pygame visualizers for Search, CSPs, Optimization, and Adversarial Games. - ob22a/ai-lab

GitHub - ob22a/ai-lab: A highly modular Object-Oriented AI framework and interactive lab implemented in Python. Features automated benchmarking, pluggable solver engines, and interactive Pygame visualizers for Search, CSPs, Optimization, and Adversarial Games. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
ob22a
/
ai-lab
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
102 Commits 102 Commits benchmarks benchmarks core core csp csp demo demo domains domains games games pdbs pdbs reports reports results results screenshots screenshots search search tests tests utils utils visualization visualization .gitignore .gitignore LICENSE LICENSE README.md README.md main.py main.py requirements.txt requirements.txt View all files Repository files navigation
🚀 Classical AI Algorithms Framework & Interactive Lab
A state-of-the-art, highly modular Object-Oriented AI framework implemented from scratch in Python. Crafted specifically for Visual Learners , researchers, and developers, this repository provides unified abstractions, pluggable solver engines, 100% transparent visualizers, automated benchmarking pipelines, and an interactive CLI Launcher ( main.py ) across a vast array of artificial intelligence paradigms.
# Install dependencies
pip install pygame matplotlib numpy
# Launch the interactive CLI menu
python main.py
# Or run any demo directly
python -m demo.maze --algo AStar --vis
python -m demo.local_search_tsp --algo GeneticAlgorithm --vis
python -m demo.crazy_demo
🎮 Unified Interactive Launcher Hub ( main.py )
Launch the interactive CLI menu to browse and run all 18 supported AI demos :
python main.py
Features of the Launcher Hub:
Categorized Menu : Browse demos across Search , Optimization , CSP , and Adversarial Games categories.
Algorithm Selector : Each demo supports CLI flags to pick algorithms (A*, BFS, DFS, UCS, IDA*, Hill Climbing, Simulated Annealing, Genetic Algorithm, Backtracking + MRV/MAC, Minimax, AlphaBeta, MCTS, IS-MCTS).
Game Modes : Human vs AI, AI vs AI, Human vs Human.
Visualizer Support : All demos with --vis flag launch interactive Pygame visualizers.
Dynamic Window Resizing : Full pygame.RESIZABLE support across all visualizers.
📸 Interactive Visualizer Gallery & Screenshot Showcase
🧩 Constraint Satisfaction & Graph Decomposition
Tree Decomposition (Junction Tree & Separators)
Cycle Cutset Conditioning (Acyclic Tree Subproblem)
Sudoku CSP (Backtracking + MRV + MAC)
N-Queens CSP Symmetry Breaking
🧭 Graph & Pathfinding Search
Maze A* Search (Manhattan / Euclidean)
Online Maze Search (LRTA* Real-Time Learning)
8-Puzzle Sliding Tile Search (Disjoint PDBs)
Romanian Map City Routing
Sokoban Box Pushing Search
Sokoban Solved State
📈 Continuous Optimization & Population Solvers
TSP Genetic Algorithm (Elite Population Inspector)
TSP Simulated Annealing (Distance Progression Curve)
N-Queens Genetic Algorithm (Top Chromosomes)
N-Queens Local Beam Search (k Parallel Beams)
N-Queens Simulated Annealing
N-Queens Hill Climbing
🎮 Adversarial Game Theory & Card Games
Crazy Card Game (Information Set MCTS vs Obssa's Heuristic)
Othello / Reversi (Alpha-Beta Pruning)
Connect Four (Alpha-Beta Search)
Checkers (Minimax & Alpha-Beta Search)
Crazy Card Game (Information Set MCTS)
Tic-Tac-Toe (Minimax Search)
🛠️ Educational Modules, Standalone Demos & CLI Execution
This Lab is built for hands-on learning, experimentation, and research. Run any demo directly from your terminal:
# 1. Graph & Pathfinding Search
python -m demo.maze --algo AStar --vis
python -m demo.maze --algo IGBFS --vis
python -m demo.n-puzzle --size 4 --algo AStar --vis
python -m demo.romanian_map_demo --start Arad --goal Bucharest --algo AStar --vis
python -m demo.sokoban_demo --vis
# 2. Local Search & Continuous Optimization
python -m demo.local_search_tsp --algo GeneticAlgorithm --vis
python -m demo.local_search_tsp --algo LocalBeamSearch --vis
python -m demo.local_search_nqueens --algo GeneticAlgorithm --vis
# 3. Constraint Satisfaction Problems (CSP)
python -m demo.csp_tree_decomposition --vis
python -m demo.csp_cycle_cutset --vis
python -m demo.csp_sudoku --difficulty hard --inference mac --vis
python -m demo.csp_map_coloring --vis
python -m demo.csp_cryptarithmetic --vis
# 4. Board Games & Imperfect Information Card Games
python -m demo.games_demo --game othello --p1 human --p2 alphabeta --vis
python -m demo.games_demo --game connect_four --p1 mcts --p2 random --vis
python -m demo.crazy_demo
🏗 Architecture & Design System
The core design philosophy of this Lab revolves around clean separation of concerns between Domains (State Representations) and Solvers (Algorithms) .
classDiagram
class SearchProblem {
+start
+goal
+get_actions(state)
+get_result(state, action)
+get_cost(state, action, next_state)
+heuristic(state)
}
class OptimizationProblem {
+initial_state
+value(state)
+get_all_neighbors(state)
+get_random_neighbor(state)
+crossover(state1, state2)
+mutate(state)
}
class CSPProblem {
+variables
+domains
+constraints
+add_constraint(constraint)
}
class GameState {
+current_player
+get_legal_actions()
+apply_action(action)
+is_terminal()
+get_utility(player)
}
class MazeSearchProblem
class NPuzzleProblem
class RomanianMapProblem
class WordLadderProblem
class SokobanProblem
class VacuumWorldProblem
class NQueensProblem
class TSPProblem
class MapColoringCSP
class NQueensCSP
class SudokuCSP
class CryptarithmeticCSP
class TimetablingCSP
class TicTacToeState
class ConnectFourState
class CheckersState
class OthelloState
class CrazyState
SearchProblem <|-- MazeSearchProblem
SearchProblem <|-- NPuzzleProblem
SearchProblem <|-- RomanianMapProblem
SearchProblem <|-- WordLadderProblem
SearchProblem <|-- SokobanProblem
SearchProblem <|-- VacuumWorldProblem
OptimizationProblem <|-- NQueensProblem
OptimizationProblem <|-- TSPProblem
CSPProblem <|-- MapColoringCSP
CSPProblem <|-- NQueensCSP
CSPProblem <|-- SudokuCSP
CSPProblem <|-- CryptarithmeticCSP
CSPProblem <|-- TimetablingCSP
GameState <|-- TicTacToeState
GameState <|-- ConnectFourState
GameState <|-- CheckersState
GameState <|-- OthelloState
GameState <|-- CrazyState
Loading
📊 Performance Benchmarks & Report Generation
Benchmark evaluation results and performance comparison reports are saved directly in markdown and high-resolution chart format.
# Run all benchmarks (30 iterations per algorithm)
python -m benchmarks.run_all_benchmarks --runs 30
# Run individual benchmark suites with filters
python -m benchmarks.search_benchmark --runs 10 --domains 8pzl --algos " A*,IDA* "
python -m benchmarks.csp_benchmark --runs 10 --algos " BT+MAC,BT+MRV "
python -m benchmarks.game_benchmark --runs 5 --games tic_tac_toe
python -m benchmarks.local_search_benchmark --runs 10 --domains tsp
# Use --reset to clear existing CSV data before writing
python -m benchmarks.search_benchmark --runs 30 --reset
Generating Reports
# Generate performance charts (saved to reports/figures/)
python -m benchmarks.generate_report
# Generate markdown summary tables (updates reports/benchmark_report.md)
python -m utils.generate_markdown_tables
Reports : See reports/benchmark_report.md and reports/comparison.md .
CSV Results : Raw per-run data in results/*.csv .
Charts : High-resolution figures in reports/figures/ .
🤝 Open Source Contribution Guidelines
We welcome open-source contributions! Adding a new search algorithm, heuristic, or domain is straightforward:
Add a New Search Algorithm : Inherit from SearchAlgorithm in search/SearchAlgorithm.py and implement search_step() .
Add a New CSP Heuristic / Inference : Implement a function receiving (csp, assignment) inside csp/heuristics/ or csp/inference/ .
Add a New Game Domain : Inherit from GameState in games/GameState.py and define get_legal_actions() , apply_action() , and is_terminal() .
Add a Custom Visualizer : Create a Pygame visualizer class in visualization/ with standardized HUD controls ( SPACE auto-play, + / - speed, LEFT / RIGHT step, R restart).
Licensed under the MIT License. Developed for educational research, visual learning, and advanced AI algorithm exploration.
A highly modular Object-Oriented AI framework and interactive lab implemented in Python. Features automated benchmarking, pluggable solver engines, and interactive Pygame visualizers for Search, CSPs, Optimization, and Adversarial Games.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
