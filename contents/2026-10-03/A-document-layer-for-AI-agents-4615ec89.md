---
source: "https://www.claix.dev"
hn_url: "https://news.ycombinator.com/item?id=49942838"
title: "A document layer for AI agents"
article_title: "Claix AI | Document Processing and Context API"
image: ""
author: "gael_dev"
captured_at: "2026-10-03T10:14:03Z"
capture_tool: "hn-digest"
hn_id: 49942838
score: 1
comments: 0
posted_at: "2026-10-03T10:02:26Z"
tags:
  - hacker-news
---

# A document layer for AI agents

- HN: [49942838](https://news.ycombinator.com/item?id=49942838)
- Source: [www.claix.dev](https://www.claix.dev)
- Score: 1
- Comments: 0
- Posted: 2026-10-03T10:02:26Z

## Translation

Title: A document layer for AI agents
Article title: Claix AI | Document Processing and Context API
Description: AI document and data infrastructure for ai systems. Turn unstructured data into structured, source-traceable JSON and persistent knowledge. Pay per use or BYOK.

Article text:
Claix AI | Document Processing and Context API { } [ ] JSON API 01 10 => <> Claix Docs Documentation Pricing ES Español Sign in Sign up Turn Unstructured Data Into Structured Intelligence
Excel / CSV → JSON POST /api/excel-json
Txt / HTML / XML → JSON POST /api/txt-json
Audio → JSON POST /api/audio-json
Claix turns PDFs, Excel files, and images into document context for developers, backends, automations, and agents—accessible via REST, A2A (Agent-to-Agent), or MCP. Query a single file by document_id, or group them under a space_id to compare, sum, and connect information across multiple documents at once.
Single Document Multi-Documents Invoice
POST /document-context/{document_id}
agent_data.json agent_data Analyzing… Data queries Cross-document or single-document query Permanent, temporary, or no document storage. YOU CHOOSE Integrate free into your flow in 5 minutes Capabilities
Your customer sees your brand, logo, and product. Claix never appears in the UI. Behind the scenes, our intelligent AI mapping engine does the heavy lifting.
REST for backends and microservices, A2A (Agent-to-Agent) when you orchestrate autonomous agents, and MCP for IDEs. No mandatory SDKs—integrate into your product stack, CRM, ERP, automation, or agents via HTTP, JSON-RPC, or MCP tools.
MCP server for IDEs and agents
Connect Claix to Cursor, Claude Desktop, or Windsurf in one click. Process PDFs and spreadsheets from your IDE—or from the agents and tools you already run—via our MCP protocol.
Claix is also a native A2A agent: peers discover it via Agent Card and invoke 20+ extraction and document-query skills over JSON-RPC. Prefer REST for your backend? It stays available unchanged.
Secure infrastructure in Stockholm
Data hosted in the EU, with encryption in transit and at rest, API Key authentication and multi-tenant isolation.
We manage the AI engine, semantic mapping, and scaling. You focus on selling your product—not maintaining data pipelines.
Full privacy control. Zero AI training.
You decide how much we remember. Neither Claix nor our providers use your confidential documents to train models. Keep absolute control over your data by choosing instant deletion (total destruction after processing), temporary mode (memory retained per session), or persistent mode (encrypted context exclusive to your organization).
Drop spaghetti code and unstable prompts. Create and manage extraction structures via API, all centralized under a unique identifier (schema_id). Your database and business logic stay clean.
A typed payload is not enough if you cannot tell a cited field from a guess. Source Tracing attaches evidence to each value, the guaranteed schema keeps the JSON contract intact, and requires_human_revision is the explicit stop signal when the document does not support the answer.
Field citations Source Tracing
Grounded extraction: each field can return { value, source } pointing to a page, cell, or fragment—not a generic badge. That is explainable AI extraction your backend, audit trail, or agent can actually check.
Typed contract Guaranteed schema
The payload still conforms to your schema. Typed JSON is the contract; Source Tracing is the evidence layer on top. Grounded AI output does not mean a looser structure.
Human-in-the-loop requires_human_revision
When anchoring fails, Claix does not invent a citation. source is requires_human_revision so human-in-the-loop AI validation can take over. That reduces AI hallucination risk instead of hiding a weak AI confidence score.
{
"invoice_total": {
"value": 1284.5,
"source": " page 1, Total Due "
},
"due_date": {
"value": "2026-10-17",
"source": " requires_human_revision "
}
} A cited total can proceed. A date without evidence is marked requires_human_revision—hallucination detection at the field, not a silent guess.
Flow architecture · REST, A2A & MCP
A frictionless pipeline for developers, backends, automations, and AI agents. Ingest heterogeneous documents, delegate extraction and context to Claix via REST, A2A, or MCP, and operate with typed data and multi-file queries.
Step 01 · Onboarding and instant API Key Frictionless onboarding
Free signup with 100 successful processing calls included. Generate your server API key in seconds and start integrating—no credit card and no OCR or embeddings infrastructure to deploy.
Step 02 · REST connection or A2A protocol REST integration and A2A protocol
Inject HTTP endpoints into your microservices, n8n, or Make with standard headers (x-api-key). In multi-agent architectures, you can also interact natively through the Agent-to-Agent (A2A) protocol with Agent Card discovery and task delegation.
Step 03 · Deterministic extraction against schema Typed extraction against schema
Claix processes scanned PDFs, Excel workbooks, Word files, or images with multimodal vision. Data is transformed deterministically into the requested JSON Schema. Fields without evidence stay null—reducing AI hallucination risk instead of inventing a number.
Step 04 · Document memory and Knowledge Spaces Persistence and Knowledge Spaces
Associate processed documents with a shared space_id. Your backend, automation, or agent can query individual files by document_id or run cross-document queries across the whole space to compare, sum, and reconcile data across multiple files in a single call.
Step 05 · Scale without tech debt Scaling and multi-tenant isolation
Strict logical isolation by space and account to guarantee privacy and GDPR compliance. Cut prompt token consumption by more than 80%, drop fragile RAG pipeline maintenance, and ship to production faster—with or without agents.
200 OK → charged 401 / 403 → free 400 / 422 → free 5xx → free PAY PER USE
Per successful call. Scale when you need to.
100 successful calls free — plenty of room to integrate Claix into your stack without spending a cent.
Excel PDF Doc Img Txt Audio Excel → JSON
Up to 5 questions per call, €0.006 per question.
For teams that need a tailored rollout.
View Success Story Go to Blog FAQ
Claix is an AI document intelligence API that converts unstructured files into schema-based JSON and queryable document context for developers, backends, automations, and AI agents. It processes PDFs, Excel spreadsheets, Word documents, images, text, HTML, XML, and audio using multimodal models, so you can extract structured data and query documents on demand. Integrate via REST, native Agent-to-Agent (A2A) (POST /a2a, Agent Card at /.well-known/agent.json), or MCP from IDEs and tool frameworks.
Yes. Claix implements the Agent-to-Agent (A2A) protocol as a first-class integration surface. Peer agents discover Claix via the Agent Card at https://www.claix.dev/.well-known/agent.json, POST JSON-RPC to /a2a, invoke 20+ document skills mapped from OpenAPI, and receive structured Task responses with DataPart payloads. A2A is ideal when one autonomous agent must delegate PDF/Excel extraction, document memory, or cross-document queries to a specialized document-intelligence agent without IDE setup. REST and MCP remain available—choose A2A for agent-to-agent orchestration.
Claix exposes a Model Context Protocol (MCP) server at /mcp so IDEs (Cursor, Claude Desktop, Windsurf) and MCP clients can invoke extraction, Agent Mode, context window, and knowledge-space tools in one click. Use MCP when a developer or copilot needs interactive document tools inside an IDE; use A2A when another autonomous agent must call Claix programmatically over the Agent-to-Agent protocol.
Claix natively supports PDF (.pdf), Excel workbooks (.xlsx, .xls), Word documents (.docx, .doc), images (.png, .jpg, .jpeg, .webp), and plain text files (.txt, .html, .md). Each format is processed with visual document understanding to return typed, validated JSON conforming to your custom schema.
To convert an Excel file, send a POST request to /api/excel-json with your file and target JSON schema. Claix interprets multi-sheet workbooks, merged headers, and relational tables, returning clean JSON ready for ingestion.
Yes. Claix provides dedicated endpoints for each format: /api/pdf-json for complex or scanned PDFs, /api/doc-json for Word files, /api/img-json for images and receipts, and /api/audio-json for spoken audio (MP3, WAV, M4A, OGG). Every endpoint outputs valid JSON according to your schema.
Claix is designed for software developers, AI agent builders, SaaS engineering teams, and automation specialists using platforms like n8n or Make. It eliminates the need to build custom OCR, chunking, and vector retrieval pipelines.
Claix offers two native modes. With Claix Managed AI you pay per successful document (from €0.10 for images to €0.20 for audio) plus €0.03 per follow-up question, with 100 free successful processing runs on new accounts and no mandatory monthly subscription for standard extraction. With Bring Your Own Key (BYOK) there is no additional Claix document-processing fee: you connect your own OpenAI, Gemini, Claude, or Grok key and pay only that provider’s token usage while Claix still handles schemas, structured JSON, document context, and automation-ready responses.
Yes. BYOK is a native Claix option: connect your own supported AI provider API key, choose a compatible model, and process PDFs, Excel, Word, images, and text through Claix with no additional Claix processing fee. You remain responsible for your provider’s token charges. Your Claix API key still authenticates the workspace. Managed AI remains available when you prefer Claix-managed rates and defaults.
Source Tracing is optional field-level citation: each extracted value can include a source anchored to the document. When that check fails, source is requires_human_revision so a person can review the field. It reduces AI hallucination risk; it does not claim to eliminate it.
No. Claix only charges for successful HTTP 200 processing requests that return valid extracted data. Failed API calls due to invalid parameters or system errors do not consume your quota or balance.
Claix Agent Mode (/agent/*-json) is an extraction mode where an AI agent interprets ambiguous documents using natural language instructions instead of a rigid schema, allowing dynamic extraction based on context.
Normal extraction enforces strict, deterministic schema conformance for predictable pipelines. Agent Mode allows natural language reasoning over document contents, ideal for unstructured contracts and variable layouts.
When you process a document with context enabled, Claix returns a document_id. You can send questions to POST /document-context/{document_id} to query that specific file without re-uploading or reprocessing.
The context window is a temporary retention layer that keeps a processed document queryable for follow-up questions without permanent storage, enabling multi-turn agent interactions before expiring automatically.
Knowledge Spaces allow you to group multiple persistent documents under a single space_id. By sending questions to POST /space-context/{space_id}, your backend, automation, or AI agent can query, compare, calculate totals, and cross-reference data across all active documents in that space simultaneously.
Yes. Claix is fully headless and integrates server-to-server via API keys, making it compatible with custom SaaS backends and automation platforms like n8n or Make.
© 2026 Claix — Semantic normalization engine · Stockholm, EU

## Original Extract

AI document and data infrastructure for ai systems. Turn unstructured data into structured, source-traceable JSON and persistent knowledge. Pay per use or BYOK.

Claix AI | Document Processing and Context API { } [ ] JSON API 01 10 => <> Claix Docs Documentation Pricing ES Español Sign in Sign up Turn Unstructured Data Into Structured Intelligence
Excel / CSV → JSON POST /api/excel-json
Txt / HTML / XML → JSON POST /api/txt-json
Audio → JSON POST /api/audio-json
Claix turns PDFs, Excel files, and images into document context for developers, backends, automations, and agents—accessible via REST, A2A (Agent-to-Agent), or MCP. Query a single file by document_id, or group them under a space_id to compare, sum, and connect information across multiple documents at once.
Single Document Multi-Documents Invoice
POST /document-context/{document_id}
agent_data.json agent_data Analyzing… Data queries Cross-document or single-document query Permanent, temporary, or no document storage. YOU CHOOSE Integrate free into your flow in 5 minutes Capabilities
Your customer sees your brand, logo, and product. Claix never appears in the UI. Behind the scenes, our intelligent AI mapping engine does the heavy lifting.
REST for backends and microservices, A2A (Agent-to-Agent) when you orchestrate autonomous agents, and MCP for IDEs. No mandatory SDKs—integrate into your product stack, CRM, ERP, automation, or agents via HTTP, JSON-RPC, or MCP tools.
MCP server for IDEs and agents
Connect Claix to Cursor, Claude Desktop, or Windsurf in one click. Process PDFs and spreadsheets from your IDE—or from the agents and tools you already run—via our MCP protocol.
Claix is also a native A2A agent: peers discover it via Agent Card and invoke 20+ extraction and document-query skills over JSON-RPC. Prefer REST for your backend? It stays available unchanged.
Secure infrastructure in Stockholm
Data hosted in the EU, with encryption in transit and at rest, API Key authentication and multi-tenant isolation.
We manage the AI engine, semantic mapping, and scaling. You focus on selling your product—not maintaining data pipelines.
Full privacy control. Zero AI training.
You decide how much we remember. Neither Claix nor our providers use your confidential documents to train models. Keep absolute control over your data by choosing instant deletion (total destruction after processing), temporary mode (memory retained per session), or persistent mode (encrypted context exclusive to your organization).
Drop spaghetti code and unstable prompts. Create and manage extraction structures via API, all centralized under a unique identifier (schema_id). Your database and business logic stay clean.
A typed payload is not enough if you cannot tell a cited field from a guess. Source Tracing attaches evidence to each value, the guaranteed schema keeps the JSON contract intact, and requires_human_revision is the explicit stop signal when the document does not support the answer.
Field citations Source Tracing
Grounded extraction: each field can return { value, source } pointing to a page, cell, or fragment—not a generic badge. That is explainable AI extraction your backend, audit trail, or agent can actually check.
Typed contract Guaranteed schema
The payload still conforms to your schema. Typed JSON is the contract; Source Tracing is the evidence layer on top. Grounded AI output does not mean a looser structure.
Human-in-the-loop requires_human_revision
When anchoring fails, Claix does not invent a citation. source is requires_human_revision so human-in-the-loop AI validation can take over. That reduces AI hallucination risk instead of hiding a weak AI confidence score.
{
"invoice_total": {
"value": 1284.5,
"source": " page 1, Total Due "
},
"due_date": {
"value": "2026-10-17",
"source": " requires_human_revision "
}
} A cited total can proceed. A date without evidence is marked requires_human_revision—hallucination detection at the field, not a silent guess.
Flow architecture · REST, A2A & MCP
A frictionless pipeline for developers, backends, automations, and AI agents. Ingest heterogeneous documents, delegate extraction and context to Claix via REST, A2A, or MCP, and operate with typed data and multi-file queries.
Step 01 · Onboarding and instant API Key Frictionless onboarding
Free signup with 100 successful processing calls included. Generate your server API key in seconds and start integrating—no credit card and no OCR or embeddings infrastructure to deploy.
Step 02 · REST connection or A2A protocol REST integration and A2A protocol
Inject HTTP endpoints into your microservices, n8n, or Make with standard headers (x-api-key). In multi-agent architectures, you can also interact natively through the Agent-to-Agent (A2A) protocol with Agent Card discovery and task delegation.
Step 03 · Deterministic extraction against schema Typed extraction against schema
Claix processes scanned PDFs, Excel workbooks, Word files, or images with multimodal vision. Data is transformed deterministically into the requested JSON Schema. Fields without evidence stay null—reducing AI hallucination risk instead of inventing a number.
Step 04 · Document memory and Knowledge Spaces Persistence and Knowledge Spaces
Associate processed documents with a shared space_id. Your backend, automation, or agent can query individual files by document_id or run cross-document queries across the whole space to compare, sum, and reconcile data across multiple files in a single call.
Step 05 · Scale without tech debt Scaling and multi-tenant isolation
Strict logical isolation by space and account to guarantee privacy and GDPR compliance. Cut prompt token consumption by more than 80%, drop fragile RAG pipeline maintenance, and ship to production faster—with or without agents.
200 OK → charged 401 / 403 → free 400 / 422 → free 5xx → free PAY PER USE
Per successful call. Scale when you need to.
100 successful calls free — plenty of room to integrate Claix into your stack without spending a cent.
Excel PDF Doc Img Txt Audio Excel → JSON
Up to 5 questions per call, €0.006 per question.
For teams that need a tailored rollout.
View Success Story Go to Blog FAQ
Claix is an AI document intelligence API that converts unstructured files into schema-based JSON and queryable document context for developers, backends, automations, and AI agents. It processes PDFs, Excel spreadsheets, Word documents, images, text, HTML, XML, and audio using multimodal models, so you can extract structured data and query documents on demand. Integrate via REST, native Agent-to-Agent (A2A) (POST /a2a, Agent Card at /.well-known/agent.json), or MCP from IDEs and tool frameworks.
Yes. Claix implements the Agent-to-Agent (A2A) protocol as a first-class integration surface. Peer agents discover Claix via the Agent Card at https://www.claix.dev/.well-known/agent.json, POST JSON-RPC to /a2a, invoke 20+ document skills mapped from OpenAPI, and receive structured Task responses with DataPart payloads. A2A is ideal when one autonomous agent must delegate PDF/Excel extraction, document memory, or cross-document queries to a specialized document-intelligence agent without IDE setup. REST and MCP remain available—choose A2A for agent-to-agent orchestration.
Claix exposes a Model Context Protocol (MCP) server at /mcp so IDEs (Cursor, Claude Desktop, Windsurf) and MCP clients can invoke extraction, Agent Mode, context window, and knowledge-space tools in one click. Use MCP when a developer or copilot needs interactive document tools inside an IDE; use A2A when another autonomous agent must call Claix programmatically over the Agent-to-Agent protocol.
Claix natively supports PDF (.pdf), Excel workbooks (.xlsx, .xls), Word documents (.docx, .doc), images (.png, .jpg, .jpeg, .webp), and plain text files (.txt, .html, .md). Each format is processed with visual document understanding to return typed, validated JSON conforming to your custom schema.
To convert an Excel file, send a POST request to /api/excel-json with your file and target JSON schema. Claix interprets multi-sheet workbooks, merged headers, and relational tables, returning clean JSON ready for ingestion.
Yes. Claix provides dedicated endpoints for each format: /api/pdf-json for complex or scanned PDFs, /api/doc-json for Word files, /api/img-json for images and receipts, and /api/audio-json for spoken audio (MP3, WAV, M4A, OGG). Every endpoint outputs valid JSON according to your schema.
Claix is designed for software developers, AI agent builders, SaaS engineering teams, and automation specialists using platforms like n8n or Make. It eliminates the need to build custom OCR, chunking, and vector retrieval pipelines.
Claix offers two native modes. With Claix Managed AI you pay per successful document (from €0.10 for images to €0.20 for audio) plus €0.03 per follow-up question, with 100 free successful processing runs on new accounts and no mandatory monthly subscription for standard extraction. With Bring Your Own Key (BYOK) there is no additional Claix document-processing fee: you connect your own OpenAI, Gemini, Claude, or Grok key and pay only that provider’s token usage while Claix still handles schemas, structured JSON, document context, and automation-ready responses.
Yes. BYOK is a native Claix option: connect your own supported AI provider API key, choose a compatible model, and process PDFs, Excel, Word, images, and text through Claix with no additional Claix processing fee. You remain responsible for your provider’s token charges. Your Claix API key still authenticates the workspace. Managed AI remains available when you prefer Claix-managed rates and defaults.
Source Tracing is optional field-level citation: each extracted value can include a source anchored to the document. When that check fails, source is requires_human_revision so a person can review the field. It reduces AI hallucination risk; it does not claim to eliminate it.
No. Claix only charges for successful HTTP 200 processing requests that return valid extracted data. Failed API calls due to invalid parameters or system errors do not consume your quota or balance.
Claix Agent Mode (/agent/*-json) is an extraction mode where an AI agent interprets ambiguous documents using natural language instructions instead of a rigid schema, allowing dynamic extraction based on context.
Normal extraction enforces strict, deterministic schema conformance for predictable pipelines. Agent Mode allows natural language reasoning over document contents, ideal for unstructured contracts and variable layouts.
When you process a document with context enabled, Claix returns a document_id. You can send questions to POST /document-context/{document_id} to query that specific file without re-uploading or reprocessing.
The context window is a temporary retention layer that keeps a processed document queryable for follow-up questions without permanent storage, enabling multi-turn agent interactions before expiring automatically.
Knowledge Spaces allow you to group multiple persistent documents under a single space_id. By sending questions to POST /space-context/{space_id}, your backend, automation, or AI agent can query, compare, calculate totals, and cross-reference data across all active documents in that space simultaneously.
Yes. Claix is fully headless and integrates server-to-server via API keys, making it compatible with custom SaaS backends and automation platforms like n8n or Make.
© 2026 Claix — Semantic normalization engine · Stockholm, EU
