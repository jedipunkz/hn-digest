---
source: "https://github.com/just-not-google/BiNeuron"
hn_url: "https://news.ycombinator.com/item?id=49949804"
title: "What if we created a local AI tool to measure your security?"
article_title: "GitHub - just-not-google/BiNeuron: BiNeuron is a local AI assistant that detects programming languages from requests/files, selects specialized models based on hardware, and generates code/docs/analysis via configurable prompts. It extracts text from PDFs, Word, images (OCR), and other formats, tran\n[truncated]"
image: "https://repository-images.githubusercontent.com/1346179151/6ecebd36-1973-4ac8-affb-0e15342353bc"
author: "BiNeuron"
captured_at: "2026-10-04T02:26:14Z"
capture_tool: "hn-digest"
hn_id: 49949804
score: 3
comments: 1
posted_at: "2026-10-04T01:51:17Z"
tags:
  - hacker-news
---

# What if we created a local AI tool to measure your security?

- HN: [49949804](https://news.ycombinator.com/item?id=49949804)
- Source: [github.com](https://github.com/just-not-google/BiNeuron)
- Score: 3
- Comments: 1
- Posted: 2026-10-04T01:51:17Z

## Translation

Title: What if we created a local AI tool to measure your security?
Article title: GitHub - just-not-google/BiNeuron: BiNeuron is a local AI assistant that detects programming languages from requests/files, selects specialized models based on hardware, and generates code/docs/analysis via configurable prompts. It extracts text from PDFs, Word, images (OCR), and other formats, tran
[truncated]
Description: BiNeuron is a local AI assistant that detects programming languages from requests/files, selects specialized models based on hardware, and generates code/docs/analysis via configurable prompts. It extracts text from PDFs, Word, images (OCR), and other formats, translates, filters profanity, and uses
[truncated]

Article text:
GitHub - just-not-google/BiNeuron: BiNeuron is a local AI assistant that detects programming languages from requests/files, selects specialized models based on hardware, and generates code/docs/analysis via configurable prompts. It extracts text from PDFs, Word, images (OCR), and other formats, translates, filters profanity, and uses proxies if needed. Runs offline. · GitHub
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
just-not-google
/
BiNeuron
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
37 Commits 37 Commits Folders and files
.github/ workflows .github/ workflows BiNeuron BiNeuron img_files img_files static static templates templates tests tests .gitignore .gitignore DOCUMENTATION.md DOCUMENTATION.md LICENSE LICENSE README.md README.md TIPS_FOR_PROMPTS.md TIPS_FOR_PROMPTS.md app.py app.py requirements.txt requirements.txt View all files Repository files navigation
Intelligent Code Analysis and Generation Platform
BiNeuron is a sophisticated software solution that bridges the gap between human intent and machine generated code. It unifies advanced natural language processing, optical character recognition, and adaptive model selection into a single, powerful tool designed for developers, researchers, and technical teams.
At its core, BiNeuron automatically identifies the programming language of a given request, extracts content from a wide array of file formats, including images and documents, and then generates context aware, production ready code using best in class local or cloud based language models.
Programming Language Detection
Supports over 25 programming languages, including Python, Java, C/C++, C#, JavaScript, TypeScript, Go, Rust, Swift, Kotlin, Ruby, Dart, Julia, Lua, SQL, MATLAB, R, Pascal, Assembly, Fortran, F#, Ada, Zig, PHP, Shell, Scala, PowerShell, Solidity, OCaml, COBOL, and more.
Combines heuristic algorithms, proprietary keyword matching, and AI powered orchestration to achieve high detection accuracy.
Analyzes both user supplied text and the content of attached files, or even entire directories.
Extracts and translates text from common document formats: PDF, Word (DOCX), ODF, PowerPoint (PPTX), Excel (XLSX/XLS), EPUB, MOBI, and FB2.
Processes source code files in nearly all text based formats, from plain text to configuration files.
Integrates two interchangeable OCR engines (EasyOCR and DeepSeek OCR) to read text from images, with optional GPU acceleration, language list selection, and automatic splitting of large images ( crop_mode ) for improved recognition.
Reads text from websites: the built in HTML scraper ( use_websites ) converts pages to clean Markdown and merges them into the request context.
Two Stage Pipeline : The primary AI model generates the code or response. A secondary, lightweight model (e.g., Qwen2.5-Coder-1.5B) then transforms the response into a strict JSON object containing absolute file paths and full new contents.
Full Context Awareness : The JSON formatter receives the complete file context (all read files, unread file names, project root, and the primary AI’s answer) to ensure accurate path generation and content mapping.
Robust Retry Mechanism : If the JSON fails validation, the system automatically re prompts the formatter up to retries times, logging each attempt until a valid JSON is produced or the maximum retries are exhausted.
Safe, Whole File Replacements : Only whole file replacements are supported (no partial edits) to maintain consistency and safety.
Optional File Deletion : When deleting_files=True is passed to the BiNeuron constructor, the JSON formatter may return null for a file path, and the system will safely delete that file. This feature is disabled by default to prevent accidental data loss.
Automatically assesses the user’s hardware capabilities (CPU cores, frequency, RAM) and selects the optimal quantized version of the target model (ranging from IQ2 to F16) to balance speed and accuracy.
Offers a curated repository of specialised models per programming language, ensuring high quality, idiomatic code generation.
Implements multi layered accessibility to Hugging Face models, including automatic fallback to hf mirror.com, dynamic proxy selection, and support for custom proxy lists.
Fetches and verifies public proxies from GitHub raw lists, with retry mechanisms and connection health checks.
Allows scanning and processing of entire folders or mounted virtual directories.
Recursively identifies supported files, extracts their content, and incorporates it into the analysis context, which is perfect for large codebases or repositories.
Interactive File Explorer : In GUI mode, the virtual storage is displayed as a tree view. Double click any file to open it in the default system application.
Built in translation engine normalises user requests to English (or any configured target language) to ensure consistent AI interactions.
Supports both Google Translate and DeepL, with automatic fallback when network restrictions are detected.
Offline Translation : Optional ArgosTranslate integration ( local_trans=True , from_code_lang='en' ) provides fully offline translation without relying on external APIs.
Text Enhancement and Compression
Request Improvement ( improving_user_experience=True ): A local small model (Qwen) rewrites the user’s raw request into a clear, structured prompt while preserving all technical details, including file names, code fragments, and error messages.
Lossless Text Compression ( compress_text=True ): When a large amount of file context is attached, the same small model compresses it, aggressively reducing token count while preserving every fact, number, and code fragment verbatim.
Request Anonymization ( anonymize_text=True ): Before being sent to translation services or AI models, the user’s request is processed through Microsoft Presidio, which detects and replaces PII (names, emails, phone numbers, addresses, credit cards, etc.) with placeholder tokens.
Profanity Filter ( filter_for_swearing=True ): Blocks requests containing aggressive language or profanity before they reach the AI.
Chat Encryption : Optionally protect all stored conversations with a master password using Fernet (AES 128 CBC + HMAC SHA256) with PBKDF2 HMAC SHA256 key derivation (200,000 iterations). After three failed unlock attempts the entire chat database and master key are permanently wiped.
BiNeuron is engineered with a modular, separation of concerns design:
Core Engine orchestrates the entire pipeline: request parsing, language detection, model selection, and response generation.
OCR Module handles text extraction from images via two interchangeable engines: EasyOCR and DeepSeek OCR (local HF Transformers or cloud API).
Model Downloader manages downloading and caching of Hugging Face models, with built in mirror and proxy support.
JSON Formatter Module uses a lightweight model (e.g., Qwen2.5-Coder-1.5B) to convert the primary model’s response into a strict JSON object for file modifications.
File Editing Module applies JSON based file changes (whole file replacements) with error handling and retry logic. Supports optional file deletion when deleting_files=True .
Text Enhancement Module handles request improvement ( PROMPT_FOR_IMPROVEMENT ) and lossless compression ( PROMPT_FOR_COMPRESSION ) via the same lightweight model.
Translation Service provides language detection and translation utilities, with optional DeepL and ArgosTranslate integration.
Anonymization Service integrates Microsoft Presidio for PII detection and masking.
Network Layer implements proxy rotation, availability checks, and GitHub proxy fetching for circumventing restrictions.
Web Interface is a feature rich web application built with Flask (HTML/CSS/JS), covering every configurable parameter of the nine configuration groups.
The architecture emphasises reusability, fault tolerance, and performance, allowing each component to operate independently while seamlessly integrating with the others.
A modern web application built with Flask, offering:
Intuitive Chat Interface : message history, file attachments, and real time log display.
Virtual Storage Explorer : scan and navigate directories in a tree view.
Comprehensive Settings Panel : fine tune every aspect of the platform, including network, model selection, prompt mode, translator, OCR, file handling, safety, and more.
Chat Management : create, delete, download, and filter conversation history.
Live Logging : see what the AI is doing in real time (with spinner, deduplicated progress bars, and one click copy).
Master Password Encryption : enable or disable chat encryption directly from the UI.
24 Dark Themes : from Midnight Deep and Dracula’s Castle to Synthwave ’84, Matrix Terminal, and AMOLED Black; every theme is tuned for long coding sessions.
Session Resume : if the page is reloaded while a request is running, the client reconnects to the ongoing task automatically.
Multilingual UI : English, Russian, and Chinese.
Only the interface and the interaction with the interface were developed with the help of DeepSeek Coder . This includes the Flask based web UI, the JavaScript front end, and the user interaction logic. All other components, including the core engine, OCR module, model downloader, JSON formatter, file editing module, translation service, anonymization service, network layer, and overall architecture, were designed and written independently.
Model Collection : deepseek-ai/deepseek-coder
Example Model : deepseek-ai/deepseek-coder-6.7b-instruct
The platform exposes nine independent configuration groups, all available from the web UI:
The complete list of libraries used by BiNeuron (exactly as declared in requirements.txt ).
Library
Purpose
Repository
torch
Tensor computation and GPU acceleration for local models
github.com/pytorch/pytorch
transformers
Loading and running Hugging Face models locally
github.com/huggingface/transformers
huggingface_hub
Downloading and caching models from Hugging Face
github.com/huggingface/huggingface_hub
llama-cpp-python
Running GGUF models via llama.cpp bindings
github.com/abetlen/llama-cpp-python
pillow
Image loading and preprocessing for OCR and vision models
github.com/python-pillow/Pillow
OCR
Library
Purpose
Repository
easyocr
OCR engine for text extraction from images
github.com/JaidedAI/EasyOCR
deepseek-ocr
DeepSeek OCR client (local and cloud)
pypi.org/project/deepseek-ocr
Document Parsing
Library
Purpose
Repository
PyMuPDF
Reading text from PDF files
github.com/pymupdf/PyMuPDF
docx2txt
Extracting text from .docx (Microsoft Word)
github.com/ankushshah89/python-docx2txt
pptx2txt2
Extracting text from .pptx (

[truncated]

## Original Extract

BiNeuron is a local AI assistant that detects programming languages from requests/files, selects specialized models based on hardware, and generates code/docs/analysis via configurable prompts. It extracts text from PDFs, Word, images (OCR), and other formats, translates, filters profanity, and uses
[truncated]

GitHub - just-not-google/BiNeuron: BiNeuron is a local AI assistant that detects programming languages from requests/files, selects specialized models based on hardware, and generates code/docs/analysis via configurable prompts. It extracts text from PDFs, Word, images (OCR), and other formats, translates, filters profanity, and uses proxies if needed. Runs offline. · GitHub
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
just-not-google
/
BiNeuron
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
37 Commits 37 Commits Folders and files
.github/ workflows .github/ workflows BiNeuron BiNeuron img_files img_files static static templates templates tests tests .gitignore .gitignore DOCUMENTATION.md DOCUMENTATION.md LICENSE LICENSE README.md README.md TIPS_FOR_PROMPTS.md TIPS_FOR_PROMPTS.md app.py app.py requirements.txt requirements.txt View all files Repository files navigation
Intelligent Code Analysis and Generation Platform
BiNeuron is a sophisticated software solution that bridges the gap between human intent and machine generated code. It unifies advanced natural language processing, optical character recognition, and adaptive model selection into a single, powerful tool designed for developers, researchers, and technical teams.
At its core, BiNeuron automatically identifies the programming language of a given request, extracts content from a wide array of file formats, including images and documents, and then generates context aware, production ready code using best in class local or cloud based language models.
Programming Language Detection
Supports over 25 programming languages, including Python, Java, C/C++, C#, JavaScript, TypeScript, Go, Rust, Swift, Kotlin, Ruby, Dart, Julia, Lua, SQL, MATLAB, R, Pascal, Assembly, Fortran, F#, Ada, Zig, PHP, Shell, Scala, PowerShell, Solidity, OCaml, COBOL, and more.
Combines heuristic algorithms, proprietary keyword matching, and AI powered orchestration to achieve high detection accuracy.
Analyzes both user supplied text and the content of attached files, or even entire directories.
Extracts and translates text from common document formats: PDF, Word (DOCX), ODF, PowerPoint (PPTX), Excel (XLSX/XLS), EPUB, MOBI, and FB2.
Processes source code files in nearly all text based formats, from plain text to configuration files.
Integrates two interchangeable OCR engines (EasyOCR and DeepSeek OCR) to read text from images, with optional GPU acceleration, language list selection, and automatic splitting of large images ( crop_mode ) for improved recognition.
Reads text from websites: the built in HTML scraper ( use_websites ) converts pages to clean Markdown and merges them into the request context.
Two Stage Pipeline : The primary AI model generates the code or response. A secondary, lightweight model (e.g., Qwen2.5-Coder-1.5B) then transforms the response into a strict JSON object containing absolute file paths and full new contents.
Full Context Awareness : The JSON formatter receives the complete file context (all read files, unread file names, project root, and the primary AI’s answer) to ensure accurate path generation and content mapping.
Robust Retry Mechanism : If the JSON fails validation, the system automatically re prompts the formatter up to retries times, logging each attempt until a valid JSON is produced or the maximum retries are exhausted.
Safe, Whole File Replacements : Only whole file replacements are supported (no partial edits) to maintain consistency and safety.
Optional File Deletion : When deleting_files=True is passed to the BiNeuron constructor, the JSON formatter may return null for a file path, and the system will safely delete that file. This feature is disabled by default to prevent accidental data loss.
Automatically assesses the user’s hardware capabilities (CPU cores, frequency, RAM) and selects the optimal quantized version of the target model (ranging from IQ2 to F16) to balance speed and accuracy.
Offers a curated repository of specialised models per programming language, ensuring high quality, idiomatic code generation.
Implements multi layered accessibility to Hugging Face models, including automatic fallback to hf mirror.com, dynamic proxy selection, and support for custom proxy lists.
Fetches and verifies public proxies from GitHub raw lists, with retry mechanisms and connection health checks.
Allows scanning and processing of entire folders or mounted virtual directories.
Recursively identifies supported files, extracts their content, and incorporates it into the analysis context, which is perfect for large codebases or repositories.
Interactive File Explorer : In GUI mode, the virtual storage is displayed as a tree view. Double click any file to open it in the default system application.
Built in translation engine normalises user requests to English (or any configured target language) to ensure consistent AI interactions.
Supports both Google Translate and DeepL, with automatic fallback when network restrictions are detected.
Offline Translation : Optional ArgosTranslate integration ( local_trans=True , from_code_lang='en' ) provides fully offline translation without relying on external APIs.
Text Enhancement and Compression
Request Improvement ( improving_user_experience=True ): A local small model (Qwen) rewrites the user’s raw request into a clear, structured prompt while preserving all technical details, including file names, code fragments, and error messages.
Lossless Text Compression ( compress_text=True ): When a large amount of file context is attached, the same small model compresses it, aggressively reducing token count while preserving every fact, number, and code fragment verbatim.
Request Anonymization ( anonymize_text=True ): Before being sent to translation services or AI models, the user’s request is processed through Microsoft Presidio, which detects and replaces PII (names, emails, phone numbers, addresses, credit cards, etc.) with placeholder tokens.
Profanity Filter ( filter_for_swearing=True ): Blocks requests containing aggressive language or profanity before they reach the AI.
Chat Encryption : Optionally protect all stored conversations with a master password using Fernet (AES 128 CBC + HMAC SHA256) with PBKDF2 HMAC SHA256 key derivation (200,000 iterations). After three failed unlock attempts the entire chat database and master key are permanently wiped.
BiNeuron is engineered with a modular, separation of concerns design:
Core Engine orchestrates the entire pipeline: request parsing, language detection, model selection, and response generation.
OCR Module handles text extraction from images via two interchangeable engines: EasyOCR and DeepSeek OCR (local HF Transformers or cloud API).
Model Downloader manages downloading and caching of Hugging Face models, with built in mirror and proxy support.
JSON Formatter Module uses a lightweight model (e.g., Qwen2.5-Coder-1.5B) to convert the primary model’s response into a strict JSON object for file modifications.
File Editing Module applies JSON based file changes (whole file replacements) with error handling and retry logic. Supports optional file deletion when deleting_files=True .
Text Enhancement Module handles request improvement ( PROMPT_FOR_IMPROVEMENT ) and lossless compression ( PROMPT_FOR_COMPRESSION ) via the same lightweight model.
Translation Service provides language detection and translation utilities, with optional DeepL and ArgosTranslate integration.
Anonymization Service integrates Microsoft Presidio for PII detection and masking.
Network Layer implements proxy rotation, availability checks, and GitHub proxy fetching for circumventing restrictions.
Web Interface is a feature rich web application built with Flask (HTML/CSS/JS), covering every configurable parameter of the nine configuration groups.
The architecture emphasises reusability, fault tolerance, and performance, allowing each component to operate independently while seamlessly integrating with the others.
A modern web application built with Flask, offering:
Intuitive Chat Interface : message history, file attachments, and real time log display.
Virtual Storage Explorer : scan and navigate directories in a tree view.
Comprehensive Settings Panel : fine tune every aspect of the platform, including network, model selection, prompt mode, translator, OCR, file handling, safety, and more.
Chat Management : create, delete, download, and filter conversation history.
Live Logging : see what the AI is doing in real time (with spinner, deduplicated progress bars, and one click copy).
Master Password Encryption : enable or disable chat encryption directly from the UI.
24 Dark Themes : from Midnight Deep and Dracula’s Castle to Synthwave ’84, Matrix Terminal, and AMOLED Black; every theme is tuned for long coding sessions.
Session Resume : if the page is reloaded while a request is running, the client reconnects to the ongoing task automatically.
Multilingual UI : English, Russian, and Chinese.
Only the interface and the interaction with the interface were developed with the help of DeepSeek Coder . This includes the Flask based web UI, the JavaScript front end, and the user interaction logic. All other components, including the core engine, OCR module, model downloader, JSON formatter, file editing module, translation service, anonymization service, network layer, and overall architecture, were designed and written independently.
Model Collection : deepseek-ai/deepseek-coder
Example Model : deepseek-ai/deepseek-coder-6.7b-instruct
The platform exposes nine independent configuration groups, all available from the web UI:
The complete list of libraries used by BiNeuron (exactly as declared in requirements.txt ).
Library
Purpose
Repository
torch
Tensor computation and GPU acceleration for local models
github.com/pytorch/pytorch
transformers
Loading and running Hugging Face models locally
github.com/huggingface/transformers
huggingface_hub
Downloading and caching models from Hugging Face
github.com/huggingface/huggingface_hub
llama-cpp-python
Running GGUF models via llama.cpp bindings
github.com/abetlen/llama-cpp-python
pillow
Image loading and preprocessing for OCR and vision models
github.com/python-pillow/Pillow
OCR
Library
Purpose
Repository
easyocr
OCR engine for text extraction from images
github.com/JaidedAI/EasyOCR
deepseek-ocr
DeepSeek OCR client (local and cloud)
pypi.org/project/deepseek-ocr
Document Parsing
Library
Purpose
Repository
PyMuPDF
Reading text from PDF files
github.com/pymupdf/PyMuPDF
docx2txt
Extracting text from .docx (Microsoft Word)
github.com/ankushshah89/python-docx2txt
pptx2txt2
Extracting text from .pptx (

[truncated]
