---
source: "https://github.com/sujunmin/agy-ppt"
hn_url: "https://news.ycombinator.com/item?id=49977935"
title: "Agy-ppt – agent skill for PowerPoint decks without the AI-slop look"
article_title: "GitHub - sujunmin/agy-ppt: Turn reports into PowerPoint decks without the AI-slop look. An agent skill purpose-built for the Antigravity + Kiro + Codex pipeline — GPT-image renders dense Traditional Chinese slide text far better than Gemini. · GitHub"
image: "https://opengraph.githubassets.com/a9e3c4b31be542832ad05cb9a80c276b6917337625641e75f40a6a5454ff6767/sujunmin/agy-ppt"
author: "sujunmin"
captured_at: "2026-10-06T13:57:59Z"
capture_tool: "hn-digest"
hn_id: 49977935
score: 2
comments: 1
posted_at: "2026-10-06T13:12:23Z"
tags:
  - hacker-news
---

# Agy-ppt – agent skill for PowerPoint decks without the AI-slop look

- HN: [49977935](https://news.ycombinator.com/item?id=49977935)
- Source: [github.com](https://github.com/sujunmin/agy-ppt)
- Score: 2
- Comments: 1
- Posted: 2026-10-06T13:12:23Z

## Translation

Title: Agy-ppt – agent skill for PowerPoint decks without the AI-slop look
Article title: GitHub - sujunmin/agy-ppt: Turn reports into PowerPoint decks without the AI-slop look. An agent skill purpose-built for the Antigravity + Kiro + Codex pipeline — GPT-image renders dense Traditional Chinese slide text far better than Gemini. · GitHub
Description: Turn reports into PowerPoint decks without the AI-slop look. An agent skill purpose-built for the Antigravity + Kiro + Codex pipeline — GPT-image renders dense Traditional Chinese slide text far better than Gemini. - sujunmin/agy-ppt

Article text:
GitHub - sujunmin/agy-ppt: Turn reports into PowerPoint decks without the AI-slop look. An agent skill purpose-built for the Antigravity + Kiro + Codex pipeline — GPT-image renders dense Traditional Chinese slide text far better than Gemini. · GitHub
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
sujunmin
/
agy-ppt
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
87 Commits 87 Commits Folders and files
.github .github .kiro/ agents .kiro/ agents assets assets examples examples scripts/ ci scripts/ ci skills/ agy-ppt skills/ agy-ppt third_party/ pypdfium2-5.13.0 third_party/ pypdfium2-5.13.0 website website .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md README_en.md README_en.md SECURITY.md SECURITY.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md View all files Repository files navigation
agy-ppt 是為 Antigravity＋Kiro＋Codex 這條 pipeline 打造的簡報 skill：Antigravity 擔任導演，負責大綱、設計方向、內容與品質把關；Kiro 負責所有工程；Codex（GPT 圖片模型）負責把每一頁算成圖，最後組裝成含講稿的 .pptx 。
為什麼算圖不用 Gemini？實測發現：Gemini 生成圖片時能承載的繁體中文很少，字一多就糊掉、亂寫；同樣的版面，GPT 撐得住幾十個中文字。簡報正好是文字密度最高的場景，這個分工是被現實逼出來的，不是偏好。
成品是混合式 PowerPoint：以圖片保留視覺細節，並在適合時保留可編輯文字、原生圖表與可替換圖片。
使用需求 ：需要可載入 skill、且具備圖片生成能力的 agent 環境（本專案以 Codex 內建圖片生成實作與測試）。這是硬需求，沒有圖片生成的環境跑不起來。
兩種風格都保留可編輯文字、適用時的原生圖表，以及可替換圖片；並非只有單一版型。瀏覽 可重現範例 以查看四張投影片預覽、prompt、大綱、風格設定與素材來源。
需求：Python 3.11+、Git，以及已登入並可使用圖片生成工具的 agent 環境。
git clone https://github.com/sujunmin/agy-ppt.git
cd agy-ppt
python3 skills/agy-ppt/scripts/codex_ppt_runtime.py bootstrap
mkdir -p ~ /.gemini/config/skills
rsync -a --delete ./skills/agy-ppt/ ~ /.gemini/config/skills/agy-ppt/
bootstrap 只建立共用 runtime 並安裝執行所需依賴；它不會執行測試、OCR qualification 或產生範例簡報。其他 agent 可把 skills/agy-ppt/ 放到其支援的 workspace skill 位置。
載入 agy-ppt skill 後，向 AGY 說明主題、受眾、頁數與來源，例如：
請把這份報告做成 10 頁的繁體中文簡報，對象是部門主管。
預設互動流程會逐步請你確認：
修改大綱或風格會使相依的樣張核准失效；在三個核准完成前，不會產生整套簡報。
視覺風格可明確選用高密度、卡片化與圖像豐富的版面；這是 opt-in 偏好，預設仍維持平衡密度，且不會用重複圖片、空泛卡片或未受來源支持的文字來湊數。
使用者需求 / 來源
→ 大綱核准
→ 風格核准
→ 單張樣張核准
→ 完整簡報
含 OCR 的來源會先建立可追溯證據，再交回 AGY 判斷：
PDF / 圖片 → 擷取 / OCR 證據 → 來源接地 → AGY
AGY 始終是唯一的 orchestrator 與語意判斷者；OCR 與圖片 worker 不會自行改寫內容或跳過核准階段。
Markdown、純文字、DOCX、靜態 HTML 與既有來源系統支援的公開遠端內容
多幀 TIFF 會被明確拒絕。OCR 準確度取決於來源品質、provider 與模型。
開發者的完整測試與 release qualification 指令位於上述 CI／測試文件，不屬於一般使用者安裝流程。
歡迎錯誤修正、文件、測試、簡報品質與版面改善、PowerPoint 相容性修正及新功能。小而聚焦的 PR 可直接提出；若會改變產品流程、AGY 語意權限、凍結契約或主要架構，請先開 issue／discussion。視覺變更必須檢查實際 render，單元測試通過本身不代表簡報品質已合格。詳見 貢獻指南 。
Phase 15 OCR 與來源接地管線、Phase 16 有根據的簡報與編輯品質、Phase 17 簡報策略與實際呈現品質，以及 Phase 18 交付真實度、可攜性與智慧編輯能力 baseline 均已完成並凍結。採用原生文字、形狀、圖表與圖片物件的混合式 PowerPoint 輸出，在確保佐證事實不可竄改與已核准視覺品質的前提下，提供業務關鍵欄位可編輯性與圖片替換能力；不宣稱 100% 全原生編輯、跨客戶端無縫相容或字型普遍可攜。Linux x86_64 是主要 production target，但正式部署仍需驗證 hard isolation；macOS 僅供開發/API 驗證，Windows 尚未通過 production-security qualification。OCR JSON schemas 仍為 deferred。
本專案採 MIT License ，衍生自 ningzimu/codex-ppt-skill ，並非 upstream 官方版本。第三方授權資料見 THIRD_PARTY_NOTICES.md 。
Turn reports into PowerPoint decks without the AI-slop look. An agent skill purpose-built for the Antigravity + Kiro + Codex pipeline — GPT-image renders dense Traditional Chinese slide text far better than Gemini.
sujunmin.github.io/agy-ppt/ Topics
Readme MIT license Contributing
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Turn reports into PowerPoint decks without the AI-slop look. An agent skill purpose-built for the Antigravity + Kiro + Codex pipeline — GPT-image renders dense Traditional Chinese slide text far better than Gemini. - sujunmin/agy-ppt

GitHub - sujunmin/agy-ppt: Turn reports into PowerPoint decks without the AI-slop look. An agent skill purpose-built for the Antigravity + Kiro + Codex pipeline — GPT-image renders dense Traditional Chinese slide text far better than Gemini. · GitHub
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
sujunmin
/
agy-ppt
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
87 Commits 87 Commits Folders and files
.github .github .kiro/ agents .kiro/ agents assets assets examples examples scripts/ ci scripts/ ci skills/ agy-ppt skills/ agy-ppt third_party/ pypdfium2-5.13.0 third_party/ pypdfium2-5.13.0 website website .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md README_en.md README_en.md SECURITY.md SECURITY.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md View all files Repository files navigation
agy-ppt 是為 Antigravity＋Kiro＋Codex 這條 pipeline 打造的簡報 skill：Antigravity 擔任導演，負責大綱、設計方向、內容與品質把關；Kiro 負責所有工程；Codex（GPT 圖片模型）負責把每一頁算成圖，最後組裝成含講稿的 .pptx 。
為什麼算圖不用 Gemini？實測發現：Gemini 生成圖片時能承載的繁體中文很少，字一多就糊掉、亂寫；同樣的版面，GPT 撐得住幾十個中文字。簡報正好是文字密度最高的場景，這個分工是被現實逼出來的，不是偏好。
成品是混合式 PowerPoint：以圖片保留視覺細節，並在適合時保留可編輯文字、原生圖表與可替換圖片。
使用需求 ：需要可載入 skill、且具備圖片生成能力的 agent 環境（本專案以 Codex 內建圖片生成實作與測試）。這是硬需求，沒有圖片生成的環境跑不起來。
兩種風格都保留可編輯文字、適用時的原生圖表，以及可替換圖片；並非只有單一版型。瀏覽 可重現範例 以查看四張投影片預覽、prompt、大綱、風格設定與素材來源。
需求：Python 3.11+、Git，以及已登入並可使用圖片生成工具的 agent 環境。
git clone https://github.com/sujunmin/agy-ppt.git
cd agy-ppt
python3 skills/agy-ppt/scripts/codex_ppt_runtime.py bootstrap
mkdir -p ~ /.gemini/config/skills
rsync -a --delete ./skills/agy-ppt/ ~ /.gemini/config/skills/agy-ppt/
bootstrap 只建立共用 runtime 並安裝執行所需依賴；它不會執行測試、OCR qualification 或產生範例簡報。其他 agent 可把 skills/agy-ppt/ 放到其支援的 workspace skill 位置。
載入 agy-ppt skill 後，向 AGY 說明主題、受眾、頁數與來源，例如：
請把這份報告做成 10 頁的繁體中文簡報，對象是部門主管。
預設互動流程會逐步請你確認：
修改大綱或風格會使相依的樣張核准失效；在三個核准完成前，不會產生整套簡報。
視覺風格可明確選用高密度、卡片化與圖像豐富的版面；這是 opt-in 偏好，預設仍維持平衡密度，且不會用重複圖片、空泛卡片或未受來源支持的文字來湊數。
使用者需求 / 來源
→ 大綱核准
→ 風格核准
→ 單張樣張核准
→ 完整簡報
含 OCR 的來源會先建立可追溯證據，再交回 AGY 判斷：
PDF / 圖片 → 擷取 / OCR 證據 → 來源接地 → AGY
AGY 始終是唯一的 orchestrator 與語意判斷者；OCR 與圖片 worker 不會自行改寫內容或跳過核准階段。
Markdown、純文字、DOCX、靜態 HTML 與既有來源系統支援的公開遠端內容
多幀 TIFF 會被明確拒絕。OCR 準確度取決於來源品質、provider 與模型。
開發者的完整測試與 release qualification 指令位於上述 CI／測試文件，不屬於一般使用者安裝流程。
歡迎錯誤修正、文件、測試、簡報品質與版面改善、PowerPoint 相容性修正及新功能。小而聚焦的 PR 可直接提出；若會改變產品流程、AGY 語意權限、凍結契約或主要架構，請先開 issue／discussion。視覺變更必須檢查實際 render，單元測試通過本身不代表簡報品質已合格。詳見 貢獻指南 。
Phase 15 OCR 與來源接地管線、Phase 16 有根據的簡報與編輯品質、Phase 17 簡報策略與實際呈現品質，以及 Phase 18 交付真實度、可攜性與智慧編輯能力 baseline 均已完成並凍結。採用原生文字、形狀、圖表與圖片物件的混合式 PowerPoint 輸出，在確保佐證事實不可竄改與已核准視覺品質的前提下，提供業務關鍵欄位可編輯性與圖片替換能力；不宣稱 100% 全原生編輯、跨客戶端無縫相容或字型普遍可攜。Linux x86_64 是主要 production target，但正式部署仍需驗證 hard isolation；macOS 僅供開發/API 驗證，Windows 尚未通過 production-security qualification。OCR JSON schemas 仍為 deferred。
本專案採 MIT License ，衍生自 ningzimu/codex-ppt-skill ，並非 upstream 官方版本。第三方授權資料見 THIRD_PARTY_NOTICES.md 。
Turn reports into PowerPoint decks without the AI-slop look. An agent skill purpose-built for the Antigravity + Kiro + Codex pipeline — GPT-image renders dense Traditional Chinese slide text far better than Gemini.
sujunmin.github.io/agy-ppt/ Topics
Readme MIT license Contributing
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
