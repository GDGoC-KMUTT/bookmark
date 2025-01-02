import { useMemo } from "react"
import ReactMarkdown from "react-markdown"
import { Prism as SyntaxHighlighter } from "react-syntax-highlighter"
import { materialOceanic as codeStyle } from "react-syntax-highlighter/dist/esm/styles/prism"
import rehypeKatex from "rehype-katex"
import rehypeRaw from "rehype-raw"
import remarkGfm from "remark-gfm"
import remarkMath from "remark-math"
import { Info, CircleCheck, TriangleAlert, Star} from "lucide-react";

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

export const transformUnderline = (text: string) => {
	const regex = /__(.*?)__/g;
	const parts = [];
	let lastIndex = 0;
	let match;

	while ((match = regex.exec(text)) !== null) {
		const start = match.index;
		const end = regex.lastIndex;

		if (start > lastIndex) {
			parts.push(text.slice(lastIndex, start));
		}
		parts.push(`<u>${match[1]}</u>`);
		lastIndex = end;
	}

	if (lastIndex < text.length) {
		parts.push(text.slice(lastIndex));
	}

	return parts.join(" ");
};

const MarkdownRenderer = ({ content }: { content: string }) => {
	let preprocessContent = useMemo(() => {
		let preprocessContent = transformInfo(content)
		preprocessContent = transformSuccess(preprocessContent)
		preprocessContent = transformWarning(preprocessContent)
		preprocessContent = transformTip(preprocessContent)
		preprocessContent = transformHighlight(preprocessContent)
		preprocessContent = transformUnderline(preprocessContent)
		return preprocessContent
	}, [content])

	return (
		<ReactMarkdown
			remarkPlugins={[remarkGfm, remarkMath]}
			rehypePlugins={[rehypeRaw, rehypeKatex]}
			components={{
				h1: ({ node, ...props }) => <h1 style={{ color: "#014FEB" }} {...props} />,
				h2: ({ node, ...props }) => <h2 style={{ color: "#007BFF", fontSize: "22px" }}  {...props} />,
				h3: ({ node, ...props }) => <h2 style={{ color: "#495057", fontSize: "19px" }}  {...props} />,
				h4: ({ node, ...props }) => <h2 style={{ color: "#6C757D", fontSize: "17px" }}  {...props} />,
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
						<SyntaxHighlighter
							style={codeStyle as any}
							language={match[1]}
							customStyle={{
								maxWidth: "850px",
								borderRadius: "1.5em",
								padding: "1.5em 3em",
								margin: "1.5em auto 2em auto",
								fontFamily: "'Courier New', monospace",
								overflowX: "auto",
								fontSize: "14px",
							}}
						>
							{String(children).replace(/\n$/, "")}
						</SyntaxHighlighter>
					) : (
						<code style={{ backgroundColor: "#FFAD88", padding: "0.2em 0.4em", fontWeight: "700" }} {...props}>
							{children}
						</code>
					)
				},
				table: ({ node, ...props }) => <table style={{ maxWidth: "90%", borderCollapse: "collapse", width: "100%", margin: "3em" }} {...props} />,
				th: ({ node, ...props }) => <th style={{ border: "1px solid gray", padding: "0.5em", fontSize: "18px" }} {...props} />,
				td: ({ node, ...props }) => <td style={{ border: "1px solid gray", padding: "0.5em" }} {...props} />,
				div: ({ node, className, ...props }) => {
					if (className === "info") {
						return (
							<div
								style={{
									backgroundColor: "#EDE5FE",
									borderRadius: "0.5em",
									display: "flex",
									alignItems: "stretch",
									gap: "1em",
									margin: "1em 0em",
								}}
								{...props}
							>
								<div
									style={{
										width: "8px",
										backgroundColor: "#5000F9",
										flexShrink: 0,
										borderTopLeftRadius: "0.5em",
										borderBottomLeftRadius: "0.5em",
									}}
								></div>
								{/* Icon */}
								<div style={{padding: "1.3em 0em 1em 0em" }}>
									<Info size={24} color="#5000F9" />
								</div>
								{/* Content */}
								<div style={{ fontWeight: 500, padding: "1em 0em 1em 0em"}}>{props.children}</div>
							</div>
						);
					}
					if (className === "success") {
						return (
							<div
								style={{
									backgroundColor: "#EEFBF2",
									borderRadius: "0.5em",
									display: "flex",
									alignItems: "stretch",
									gap: "1em",
									margin: "1em 0em",

								}}
								{...props}
							>
								<div
									style={{
										width: "8px",
										backgroundColor: "#54D47C",
										flexShrink: 0,
										borderTopLeftRadius: "0.5em",
										borderBottomLeftRadius: "0.5em",
									}}
								></div>
								{/* Icon */}
								<div style={{padding: "1.3em 0em 1em 0em" }}>
									<CircleCheck size={24} color="#54D47C" />
								</div>
								{/* Content */}
								<div style={{ fontWeight: 500, padding: "1em 0em 1em 0em"}}>{props.children}</div>
							</div>
						);
					}
					if (className === "warning") {
						return (
							<div
								style={{
									backgroundColor: "#F9EBEC",
									borderRadius: "0.5em",
									display: "flex",
									alignItems: "stretch",
									gap: "1em",
									margin: "1em 0em",

								}}
								{...props}
							>
								<div
									style={{
										width: "8px",
										backgroundColor: "#C14043",
										flexShrink: 0,
										borderTopLeftRadius: "0.5em",
										borderBottomLeftRadius: "0.5em",
									}}
								></div>
								{/* Icon */}
								<div style={{padding: "1.3em 0em 1em 0em" }}>
									<TriangleAlert size={24} color="#C14043" />
								</div>
								{/* Content */}
								<div style={{ fontWeight: 500, padding: "1em 0em 1em 0em"}}>{props.children}</div>
							</div>
						);
					}
					if (className === "tip") {
						return (
							<div
								style={{
									backgroundColor: "#FCF8EB",
									borderRadius: "0.5em",
									display: "flex",
									alignItems: "stretch",
									gap: "1em",
									margin: "1em 0em",

								}}
								{...props}
							>
								<div
									style={{
										width: "8px",
										backgroundColor: "#E6BD3B",
										flexShrink: 0,
										borderTopLeftRadius: "0.5em",
										borderBottomLeftRadius: "0.5em",
									}}
								></div>
								{/* Icon */}
								<div style={{padding: "1.3em 0em 1em 0em" }}>
									<Star size={24} color="#E6BD3B" />
								</div>
								{/* Content */}
								<div style={{ fontWeight: 500, padding: "1em 0em 1em 0em"}}>{props.children}</div>
							</div>
						);
					}
					return <div {...props} />;
				},

			}}
		>
			{preprocessContent}
		</ReactMarkdown>
	)
}

export default MarkdownRenderer

