import { useMemo } from "react"
import ReactMarkdown from "react-markdown"
import { Prism as SyntaxHighlighter } from "react-syntax-highlighter"
import { materialOceanic as codeStyle } from "react-syntax-highlighter/dist/esm/styles/prism"
import rehypeKatex from "rehype-katex"
import rehypeRaw from "rehype-raw"
import remarkGfm from "remark-gfm"
import remarkMath from "remark-math"

export const transformHighlight = (text: string) => {
    const regex = /==(.*?)==/g
    const parts = []
    let lastIndex = 0
    let match

    while ((match = regex.exec(text)) !== null) {
        const start = match.index
        const end = regex.lastIndex

        if (start > lastIndex) {
            parts.push(text.slice(lastIndex, start))
        }
        parts.push(`<mark>${match[1]}</mark>`)
        lastIndex = end
    }

    if (lastIndex < text.length) {
        parts.push(text.slice(lastIndex))
    }

    return parts.join(" ")
}

export const transformInfo = (text: string) => {
    const regex = /:::info([\s\S]*?):::/g
    const parts = []
    let lastIndex = 0
    let match

    while ((match = regex.exec(text)) !== null) {
        const start = match.index
        const end = regex.lastIndex

        if (start > lastIndex) {
            parts.push(text.slice(lastIndex, start))
        }

        parts.push(`<div class="info">${match[1]}</div>`)
        lastIndex = end
    }

    if (lastIndex < text.length) {
        parts.push(text.slice(lastIndex))
    }

    return parts.join(" ")
}

export const transformWarning = (text: string) => {
    const regex = /:::warning([\s\S]*?):::/g
    const parts = []
    let lastIndex = 0
    let match

    while ((match = regex.exec(text)) !== null) {
        const start = match.index
        const end = regex.lastIndex

        if (start > lastIndex) {
            parts.push(text.slice(lastIndex, start))
        }

        parts.push(`<div class="warning">${match[1]}</div>`)
        lastIndex = end
    }

    if (lastIndex < text.length) {
        parts.push(text.slice(lastIndex))
    }

    return parts.join(" ")
}

export const transformSuccess = (text: string) => {
    const regex = /:::success([\s\S]*?):::/g
    const parts = []
    let lastIndex = 0
    let match

    while ((match = regex.exec(text)) !== null) {
        const start = match.index
        const end = regex.lastIndex

        if (start > lastIndex) {
            parts.push(text.slice(lastIndex, start))
        }

        parts.push(`<div class="success">${match[1]}</div>`)
        lastIndex = end
    }

    if (lastIndex < text.length) {
        parts.push(text.slice(lastIndex))
    }

    return parts.join(" ")
}

export const transformTip = (text: string) => {
    const regex = /:::tip([\s\S]*?):::/g
    const parts = []
    let lastIndex = 0
    let match

    while ((match = regex.exec(text)) !== null) {
        const start = match.index
        const end = regex.lastIndex

        if (start > lastIndex) {
            parts.push(text.slice(lastIndex, start))
        }

        parts.push(`<div class="tip">${match[1]}</div>`)
        lastIndex = end
    }

    if (lastIndex < text.length) {
        parts.push(text.slice(lastIndex))
    }

    return parts.join(" ")
}

const MarkdownRenderer = ({ content }: { content: string }) => {
    let preprocessContent = useMemo(() => {
        let preprocessContent = transformInfo(content)
        preprocessContent = transformSuccess(preprocessContent)
        preprocessContent = transformWarning(preprocessContent)
        preprocessContent = transformTip(preprocessContent)
        preprocessContent = transformHighlight(preprocessContent)
        return preprocessContent
    }, [content])

    return (
        <ReactMarkdown
            remarkPlugins={[remarkGfm, remarkMath]}
            rehypePlugins={[rehypeRaw, rehypeKatex]}
            components={{
                h1: ({ node, ...props }) => <h1 style={{ color: "teal" }} {...props} />,
                h2: ({ node, ...props }) => <h2 style={{ color: "slateblue" }} {...props} />,
                p: ({ children, ...props }) => {
                    return (
                        <p style={{ lineHeight: "1.6" }} {...props}>
                            {children}
                        </p>
                    )
                },
                code: ({ node, className, children, ...props }) => {
                    const match = /language-(\w+)/.exec(className || "")
                    return match ? (
                        <SyntaxHighlighter style={codeStyle as any} language={match[1]}>
                            {String(children).replace(/\n$/, "")}
                        </SyntaxHighlighter>
                    ) : (
                        <code style={{ backgroundColor: "#f4f4f4", padding: "0.2em 0.4em" }} {...props}>
                            {children}
                        </code>
                    )
                },
                table: ({ node, ...props }) => <table style={{ borderCollapse: "collapse", width: "100%" }} {...props} />,
                th: ({ node, ...props }) => <th style={{ border: "1px solid black", padding: "0.5em" }} {...props} />,
                td: ({ node, ...props }) => <td style={{ border: "1px solid black", padding: "0.5em" }} {...props} />,
                div: ({ node, className, ...props }) => {
                    if (className === "info") {
                        return <div style={{ border: "1px solid blue", padding: "1em", borderRadius: "4px" }} {...props} />
                    }
                    if (className === "warning") {
                        return <div style={{ border: "1px solid orange", padding: "1em", borderRadius: "4px" }} {...props} />
                    }
                    if (className === "success") {
                        return <div style={{ border: "1px solid green", padding: "1em", borderRadius: "4px" }} {...props} />
                    }
                    if (className === "tip") {
                        return <div style={{ border: "1px solid purple", padding: "1em", borderRadius: "4px" }} {...props} />
                    }
                    return <div {...props} />
                },
            }}
        >
            {preprocessContent}
        </ReactMarkdown>
    )
}

export default MarkdownRenderer

