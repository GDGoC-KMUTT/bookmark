/**
 * Generates a fallback name from a given text by taking the first character
 * of up to two words and converting them to uppercase.
 *
 * @param text - The input text to generate the fallback name.
 * @returns The generated fallback name (e.g., initials).
 */
export const getFallbackName = (text: string): string => {
    const parts = text.split(" ")
    return parts
        .slice(0, 2)
        .map((part) => part.charAt(0).toUpperCase())
        .join("")
}
