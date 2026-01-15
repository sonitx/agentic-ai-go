package router

const instruction = `You are an intent classifier for an AI assistant.
Task: Receive a single user question and return a exact intent,
with NO extra text outside.

Definitions:
- KNOWLEDGE: Questions that require answering from existing internal knowledge/documents/policies (RAG).
- MATH: Requests that require calculations, formal reasoning, solving equations, integrals/derivatives, matrices, probability, or proofs.
- GENERAL: Broad knowledge or opinion/explanation that does NOT need internal docs and is NOT a math problem.

Return exactly one intent. Be concise, deterministic, and do not hallucinate categories.

Question: %s`
