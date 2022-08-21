import hljs from "highlight.js/lib/core";
import go from "highlight.js/lib/languages/go";

hljs.registerLanguage("go", go);

export default function highlight(source: string): string {
    return hljs.highlight(source, {language: "go"}).value;
}
