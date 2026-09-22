---
source: "https://github.com/ai-mode/ai-mode"
hn_url: "https://news.ycombinator.com/item?id=49798270"
title: "AI Mode for Emacs"
article_title: "GitHub - ai-mode/ai-mode: AI mode for Emacs · GitHub"
image: "https://opengraph.githubassets.com/79d7503a7a8d76afde7d7b3850089657678871aeb6a3e14f1a65ea7a26bea971/ai-mode/ai-mode"
author: "xvilka"
captured_at: "2026-09-22T09:05:04Z"
capture_tool: "hn-digest"
hn_id: 49798270
score: 1
comments: 0
posted_at: "2026-09-22T08:52:02Z"
tags:
  - hacker-news
---

# AI Mode for Emacs

- HN: [49798270](https://news.ycombinator.com/item?id=49798270)
- Source: [github.com](https://github.com/ai-mode/ai-mode)
- Score: 1
- Comments: 0
- Posted: 2026-09-22T08:52:02Z

## Translation

Title: AI Mode for Emacs
Article title: GitHub - ai-mode/ai-mode: AI mode for Emacs · GitHub
Description: AI mode for Emacs. Contribute to ai-mode/ai-mode development by creating an account on GitHub.

Article text:
GitHub - ai-mode/ai-mode: AI mode for Emacs · GitHub
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
ai-mode
/
ai-mode
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Latest commit
89 Commits 89 Commits Folders and files
.ai .ai media media .ai-ignore .ai-ignore .gitignore .gitignore Eask Eask LICENSE LICENSE Makefile Makefile README.md README.md ai-chat.el ai-chat.el ai-command-management.el ai-command-management.el ai-common.el ai-common.el ai-completions.el ai-completions.el ai-context-management.el ai-context-management.el ai-core.el ai-core.el ai-debug.el ai-debug.el ai-execution.el ai-execution.el ai-file-system.el ai-file-system.el ai-logging.el ai-logging.el ai-mode-adapter-api.el ai-mode-adapter-api.el ai-mode-indexing.el ai-mode-indexing.el ai-mode-line.el ai-mode-line.el ai-mode.el ai-mode.el ai-model-management.el ai-model-management.el ai-network.el ai-network.el ai-progress.el ai-progress.el ai-project.el ai-project.el ai-prompt-management-debug.el ai-prompt-management-debug.el ai-prompt-management.el ai-prompt-management.el ai-request-audit.el ai-request-audit.el ai-response-processors-debug.el ai-response-processors-debug.el ai-response-processors.el ai-response-processors.el ai-structs.el ai-structs.el ai-telemetry.el ai-telemetry.el ai-usage.el ai-usage.el ai-user-input.el ai-user-input.el ai-utils.el ai-utils.el ai.el ai.el View all files Repository files navigation
Installation
Installation via MELPA
Usage & Commands
Code Completion
System & Session Context Management
Extending AI Mode with Custom Commands
AI mode for Emacs is a comprehensive package that integrates powerful artificial intelligence capabilities directly into your Emacs workflow. This package transforms Emacs into an AI-enhanced development environment, providing intelligent assistance for coding, writing, and general text manipulation tasks. The modular architecture ensures maintainability, testability, and extensibility.
Interactive AI Chat with Buffer Binding : Directly engage with AI models for real-time problem-solving, code review, and general assistance. Create dedicated chat sessions bound to specific buffers with rich context integration, smart startup strategies, and context-based chat creation from response buffers.
Intelligent Code Completion : Receive interactive, multi-candidate code suggestions and completions. It supports real-time previews, navigation through candidates, dynamic context adjustment, and integration of user instructions.
Code Modification & Refactoring : Streamline development by leveraging AI to intelligently modify, improve, and refactor existing code.
Code Analysis & Documentation : Generate comprehensive explanations, documentation, and insightful analysis for code blocks and functions, enhancing readability and maintainability.
Automated Bug Detection & Fixing : Utilize AI to detect common issues and suggest fixes directly within your buffer.
Multi-Backend Support : Seamlessly integrate with various AI providers, including OpenAI, Anthropic, Hugging Face, and Google Generative AI, for diverse model access.
Customizable Commands : Tailor AI operations to your specific development needs by defining custom commands and integrating them seamlessly into your Emacs setup.
Advanced Context Management : Manage global, buffer-local, and temporary context pools, with extended context capabilities including project-wide indexing, memory files, and comprehensive context enrichment.
Real-time Previews & Progress Indicators : Get instant visual feedback for completions and track the progress of ongoing AI requests.
Structured Request Auditing : Detailed logging of AI requests, contexts, and responses for debugging, performance analysis, and security auditing with improved callback system and error handling.
Whether you're debugging complex code, exploring new technologies, or enhancing your documentation, AI mode provides the tools to accelerate your productivity and enhance your Emacs experience.
The easiest and recommended way to install this package is through MELPA. To configure MELPA, add the following lines to your ~/.emacs file:
( package-initialize )
( add-to-list 'package-archives '( " melpa " . " http://melpa.org/packages/ " ))
Then, run the command:
M-x package-install RET ai-mode RET
M-x package-install RET ai-mode-openai RET # If you plan to use OpenAI API
M-x package-install RET ai-mode-anthropic RET # If you plan to use Anthropic
M-x package-install RET ai-mode-google-genai RET # If you plan to use Google Generative AI
You might want to refresh the package contents if you haven't done so recently by running:
M-x package-refresh-contents RET
Installation via GIT
If you prefer to manage packages manually or use a different package manager, clone the ai-mode repository and its related backend repositories into a common directory, for example, ~/.emacs.d/plugins/ :
$ mkdir -p ~ /.emacs.d/plugins
$ cd ~ /.emacs.d/plugins
$ git clone https://github.com/ai-mode/ai-mode
$ git clone https://github.com/ai-mode/ai-mode-openai # If you plan to use OpenAI API
$ git clone https://github.com/ai-mode/ai-mode-anthropic # If you plan to use Anthropic
$ git clone https://github.com/ai-mode/ai-mode-deepseek # If you plan to use DeepSeek
$ git clone https://github.com/ai-mode/ai-mode-hf # If you plan to use Hugging Face
$ git clone https://github.com/ai-mode/ai-mode-google-genai # If you plan to use Google Generative AI
$ git clone https://github.com/ai-mode/ob-ai # If you plan to use Org Babel integration
Then, add the ~/.emacs.d/plugins/ directory to your Emacs load-path in your ~/.emacs or init.el file. For use-package , you can configure it by adding the following to your init.el (or .emacs ):
( add-to-list 'load-path ( expand-file-name " ~/.emacs.d/plugins/ " ))
; ; You can then use use-package as usual for ai-mode and its backends,
; ; as they will be discoverable on your load-path.
Configuration
AI Mode's functionality is highly customizable. While this package provides the core framework, integration with specific AI models (like OpenAI, Anthropic, etc.) is handled by separate backend packages.
To enable a specific backend, you typically load its package and then register its model-providing function with ai-mode . This makes its models available for use in ai-mode and ai-chat . Remember to configure any necessary API keys as per the backend's documentation.
Here's an example of a comprehensive use-package configuration for ai-mode , including how to enable various backends and the ai-debug package:
( use-package ai-mode
:defer t
:config ( progn
(global-ai-mode)
; ; Set global keybindings for code completion
( global-set-key ( kbd " C-<tab> " ) 'ai-completions-complete-code-at-point-with-limited-context )
( global-set-key ( kbd " C-M-<tab> " ) 'ai-completions-complete-at-point-with-full-context )
; ; Enable completions mode for programming buffers
( add-hook 'prog-mode-hook # 'ai-completions-mode )
))
; ; Example of enabling specific backends
; ; Ensure you have installed the respective packages first (e.g., ai-mode-openai, ai-mode-google-genai).
; ; OpenAI Backend
( use-package ai-mode-openai
:after (ai-mode ai-model-management)
:config ( progn
( add-to-list 'ai-model-management-providers 'ai-mode-openai--get-models )
; ; Set your OpenAI API key:
; ; (setq ai-mode-openai--api-key "YOUR_OPENAI_API_KEY")
))
; ; Google Generative AI Backend
( use-package ai-mode-google-genai
:after (ai-mode ai-model-management)
:config ( progn
( add-to-list 'ai-model-management-providers 'ai-mode-google-genai--get-models )
; ; Set your Google Generative AI API key:
; ; (setq ai-mode-google-genai--api-key "YOUR_GOOGLE_GENAI_API_KEY")
))
; ; Anthropic Backend
( use-package ai-mode-anthropic
:load-path " ~/.emacs.d/plugins/ai-mode-anthropic " ; ; Example for local install
:after (ai-mode ai-model-management)
:config ( progn
( add-to-list 'ai-model-management-providers 'ai-mode-anthropic--get-models )
; ; Set your Anthropic API key:
; ; (setq ai-mode-anthropic--api-key "YOUR_ANTHROPIC_API_KEY")
))
; ; Other backends (DeepSeek, Hugging Face) can be configured similarly.
; ; AI Debug Package
( use-package ai-debug
:after (ai-mode)
:config ( progn
; ; Configure debug features, e.g., disable content truncation for full view
( setq ai-debug-truncate-content nil )
( setq ai-debug-max-content-length 2000 )
))
Supported Backends
AI mode integrates with a variety of AI providers through dedicated backend packages. To utilize a specific backend, ensure you have installed its corresponding package and configured any necessary API keys.
AI Mode OpenAI : OpenAI GPT backend for ai-mode .
AI Mode Anthropic : Anthropic Claude backend for ai-mode .
AI Mode DeepSeek : DeepSeek backend for ai-mode .
AI Mode Hugging Face : Hugging Face models backend for ai-mode .
AI Mode Google Generative AI : Google Generative AI backend for ai-mode .
AI Mode provides a rich set of interactive commands, many of which are accessible through the ai-command-map prefix (default: C-c i ).
AI Mode offers intelligent code completion at the cursor position. You can use the primary entry point ai-completions-complete-code-at-point to initiate a completion session, which can be further refined with context-specific commands.
When a completion session is active, the following interactive commands are available to navigate and manage candidates:
The ai-chat package provides a full-featured interactive chat interface with AI models, built on Emacs's comint-mode for a familiar and extensible experience. It supports asynchronous interaction with various AI backends, allowing you to ask questions, review code, or get general assistance in real-time.
Start a chat: Invoke M-x ai-chat or use the default keybinding C-c i c c . This opens a dedicated *ai-chat* buffer.
Send messages: Type your message at the AI> prompt and press C-return to send it to the AI.
Manage Sessions: Chat sessions can be automatically or manually saved and loaded for continuity.
Progress Indicators: Visual indicators are displayed in the mode line to show when an AI request is in progress.
ai-chat offers several customizable variables to tailor the chat experience. These are listed in the Customization Options section.
Keybindings (within the *ai-chat* buffer):
AI Mode supports advanced buffer-bound chat sessions that create dedicated chat buffers tied to specific source files. This feature enables context-aware conversations with rich startup context and seamless integration between your code and AI discussions.
Per-Buffer Chat Sessions : Each buffer can have one or more dedicated chat sessions
Multiple Session Support : Create multiple parallel ch

[truncated]

## Original Extract

AI mode for Emacs. Contribute to ai-mode/ai-mode development by creating an account on GitHub.

GitHub - ai-mode/ai-mode: AI mode for Emacs · GitHub
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
ai-mode
/
ai-mode
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Latest commit
89 Commits 89 Commits Folders and files
.ai .ai media media .ai-ignore .ai-ignore .gitignore .gitignore Eask Eask LICENSE LICENSE Makefile Makefile README.md README.md ai-chat.el ai-chat.el ai-command-management.el ai-command-management.el ai-common.el ai-common.el ai-completions.el ai-completions.el ai-context-management.el ai-context-management.el ai-core.el ai-core.el ai-debug.el ai-debug.el ai-execution.el ai-execution.el ai-file-system.el ai-file-system.el ai-logging.el ai-logging.el ai-mode-adapter-api.el ai-mode-adapter-api.el ai-mode-indexing.el ai-mode-indexing.el ai-mode-line.el ai-mode-line.el ai-mode.el ai-mode.el ai-model-management.el ai-model-management.el ai-network.el ai-network.el ai-progress.el ai-progress.el ai-project.el ai-project.el ai-prompt-management-debug.el ai-prompt-management-debug.el ai-prompt-management.el ai-prompt-management.el ai-request-audit.el ai-request-audit.el ai-response-processors-debug.el ai-response-processors-debug.el ai-response-processors.el ai-response-processors.el ai-structs.el ai-structs.el ai-telemetry.el ai-telemetry.el ai-usage.el ai-usage.el ai-user-input.el ai-user-input.el ai-utils.el ai-utils.el ai.el ai.el View all files Repository files navigation
Installation
Installation via MELPA
Usage & Commands
Code Completion
System & Session Context Management
Extending AI Mode with Custom Commands
AI mode for Emacs is a comprehensive package that integrates powerful artificial intelligence capabilities directly into your Emacs workflow. This package transforms Emacs into an AI-enhanced development environment, providing intelligent assistance for coding, writing, and general text manipulation tasks. The modular architecture ensures maintainability, testability, and extensibility.
Interactive AI Chat with Buffer Binding : Directly engage with AI models for real-time problem-solving, code review, and general assistance. Create dedicated chat sessions bound to specific buffers with rich context integration, smart startup strategies, and context-based chat creation from response buffers.
Intelligent Code Completion : Receive interactive, multi-candidate code suggestions and completions. It supports real-time previews, navigation through candidates, dynamic context adjustment, and integration of user instructions.
Code Modification & Refactoring : Streamline development by leveraging AI to intelligently modify, improve, and refactor existing code.
Code Analysis & Documentation : Generate comprehensive explanations, documentation, and insightful analysis for code blocks and functions, enhancing readability and maintainability.
Automated Bug Detection & Fixing : Utilize AI to detect common issues and suggest fixes directly within your buffer.
Multi-Backend Support : Seamlessly integrate with various AI providers, including OpenAI, Anthropic, Hugging Face, and Google Generative AI, for diverse model access.
Customizable Commands : Tailor AI operations to your specific development needs by defining custom commands and integrating them seamlessly into your Emacs setup.
Advanced Context Management : Manage global, buffer-local, and temporary context pools, with extended context capabilities including project-wide indexing, memory files, and comprehensive context enrichment.
Real-time Previews & Progress Indicators : Get instant visual feedback for completions and track the progress of ongoing AI requests.
Structured Request Auditing : Detailed logging of AI requests, contexts, and responses for debugging, performance analysis, and security auditing with improved callback system and error handling.
Whether you're debugging complex code, exploring new technologies, or enhancing your documentation, AI mode provides the tools to accelerate your productivity and enhance your Emacs experience.
The easiest and recommended way to install this package is through MELPA. To configure MELPA, add the following lines to your ~/.emacs file:
( package-initialize )
( add-to-list 'package-archives '( " melpa " . " http://melpa.org/packages/ " ))
Then, run the command:
M-x package-install RET ai-mode RET
M-x package-install RET ai-mode-openai RET # If you plan to use OpenAI API
M-x package-install RET ai-mode-anthropic RET # If you plan to use Anthropic
M-x package-install RET ai-mode-google-genai RET # If you plan to use Google Generative AI
You might want to refresh the package contents if you haven't done so recently by running:
M-x package-refresh-contents RET
Installation via GIT
If you prefer to manage packages manually or use a different package manager, clone the ai-mode repository and its related backend repositories into a common directory, for example, ~/.emacs.d/plugins/ :
$ mkdir -p ~ /.emacs.d/plugins
$ cd ~ /.emacs.d/plugins
$ git clone https://github.com/ai-mode/ai-mode
$ git clone https://github.com/ai-mode/ai-mode-openai # If you plan to use OpenAI API
$ git clone https://github.com/ai-mode/ai-mode-anthropic # If you plan to use Anthropic
$ git clone https://github.com/ai-mode/ai-mode-deepseek # If you plan to use DeepSeek
$ git clone https://github.com/ai-mode/ai-mode-hf # If you plan to use Hugging Face
$ git clone https://github.com/ai-mode/ai-mode-google-genai # If you plan to use Google Generative AI
$ git clone https://github.com/ai-mode/ob-ai # If you plan to use Org Babel integration
Then, add the ~/.emacs.d/plugins/ directory to your Emacs load-path in your ~/.emacs or init.el file. For use-package , you can configure it by adding the following to your init.el (or .emacs ):
( add-to-list 'load-path ( expand-file-name " ~/.emacs.d/plugins/ " ))
; ; You can then use use-package as usual for ai-mode and its backends,
; ; as they will be discoverable on your load-path.
Configuration
AI Mode's functionality is highly customizable. While this package provides the core framework, integration with specific AI models (like OpenAI, Anthropic, etc.) is handled by separate backend packages.
To enable a specific backend, you typically load its package and then register its model-providing function with ai-mode . This makes its models available for use in ai-mode and ai-chat . Remember to configure any necessary API keys as per the backend's documentation.
Here's an example of a comprehensive use-package configuration for ai-mode , including how to enable various backends and the ai-debug package:
( use-package ai-mode
:defer t
:config ( progn
(global-ai-mode)
; ; Set global keybindings for code completion
( global-set-key ( kbd " C-<tab> " ) 'ai-completions-complete-code-at-point-with-limited-context )
( global-set-key ( kbd " C-M-<tab> " ) 'ai-completions-complete-at-point-with-full-context )
; ; Enable completions mode for programming buffers
( add-hook 'prog-mode-hook # 'ai-completions-mode )
))
; ; Example of enabling specific backends
; ; Ensure you have installed the respective packages first (e.g., ai-mode-openai, ai-mode-google-genai).
; ; OpenAI Backend
( use-package ai-mode-openai
:after (ai-mode ai-model-management)
:config ( progn
( add-to-list 'ai-model-management-providers 'ai-mode-openai--get-models )
; ; Set your OpenAI API key:
; ; (setq ai-mode-openai--api-key "YOUR_OPENAI_API_KEY")
))
; ; Google Generative AI Backend
( use-package ai-mode-google-genai
:after (ai-mode ai-model-management)
:config ( progn
( add-to-list 'ai-model-management-providers 'ai-mode-google-genai--get-models )
; ; Set your Google Generative AI API key:
; ; (setq ai-mode-google-genai--api-key "YOUR_GOOGLE_GENAI_API_KEY")
))
; ; Anthropic Backend
( use-package ai-mode-anthropic
:load-path " ~/.emacs.d/plugins/ai-mode-anthropic " ; ; Example for local install
:after (ai-mode ai-model-management)
:config ( progn
( add-to-list 'ai-model-management-providers 'ai-mode-anthropic--get-models )
; ; Set your Anthropic API key:
; ; (setq ai-mode-anthropic--api-key "YOUR_ANTHROPIC_API_KEY")
))
; ; Other backends (DeepSeek, Hugging Face) can be configured similarly.
; ; AI Debug Package
( use-package ai-debug
:after (ai-mode)
:config ( progn
; ; Configure debug features, e.g., disable content truncation for full view
( setq ai-debug-truncate-content nil )
( setq ai-debug-max-content-length 2000 )
))
Supported Backends
AI mode integrates with a variety of AI providers through dedicated backend packages. To utilize a specific backend, ensure you have installed its corresponding package and configured any necessary API keys.
AI Mode OpenAI : OpenAI GPT backend for ai-mode .
AI Mode Anthropic : Anthropic Claude backend for ai-mode .
AI Mode DeepSeek : DeepSeek backend for ai-mode .
AI Mode Hugging Face : Hugging Face models backend for ai-mode .
AI Mode Google Generative AI : Google Generative AI backend for ai-mode .
AI Mode provides a rich set of interactive commands, many of which are accessible through the ai-command-map prefix (default: C-c i ).
AI Mode offers intelligent code completion at the cursor position. You can use the primary entry point ai-completions-complete-code-at-point to initiate a completion session, which can be further refined with context-specific commands.
When a completion session is active, the following interactive commands are available to navigate and manage candidates:
The ai-chat package provides a full-featured interactive chat interface with AI models, built on Emacs's comint-mode for a familiar and extensible experience. It supports asynchronous interaction with various AI backends, allowing you to ask questions, review code, or get general assistance in real-time.
Start a chat: Invoke M-x ai-chat or use the default keybinding C-c i c c . This opens a dedicated *ai-chat* buffer.
Send messages: Type your message at the AI> prompt and press C-return to send it to the AI.
Manage Sessions: Chat sessions can be automatically or manually saved and loaded for continuity.
Progress Indicators: Visual indicators are displayed in the mode line to show when an AI request is in progress.
ai-chat offers several customizable variables to tailor the chat experience. These are listed in the Customization Options section.
Keybindings (within the *ai-chat* buffer):
AI Mode supports advanced buffer-bound chat sessions that create dedicated chat buffers tied to specific source files. This feature enables context-aware conversations with rich startup context and seamless integration between your code and AI discussions.
Per-Buffer Chat Sessions : Each buffer can have one or more dedicated chat sessions
Multiple Session Support : Create multiple parallel ch

[truncated]
