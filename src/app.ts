import "./wasm/wasm_exec"

import highlight from "./highlight";
import {Options} from "./convert";

const $input = document.getElementById("input");
const $output = document.getElementById("output");
const $sample = document.getElementById("sample");
const $inline = document.getElementById("inline") as HTMLInputElement;
const $withJSON = document.getElementById("with-json-tags") as HTMLInputElement;
const $compact = document.getElementById("compact") as HTMLInputElement;

const options = new Options($inline.checked, $withJSON.checked, $compact.checked);

function doConversion() {
    const result = globalThis.xmlDataToGoTypeCode($input.innerText.trim(), options.inline, options.compact, options.withJSON);

    if (result !== "") {
        $output.innerHTML = highlight(result);
    } else {
        $output.innerHTML = "";
    }
}

function xmlToGoWasmURL() {
    if (location.pathname.startsWith("/home/")) {
        return "http://localhost:8080/xml-to-go.wasm";
    }

    return "https://xml-to-go.github.io/static/js/wasm/xml-to-go.wasm";
}

$input.addEventListener("keyup", doConversion);

$inline.addEventListener("change", function () {
    options.inline = $inline.checked;

    doConversion();
});

$withJSON.addEventListener("change", function () {
    options.withJSON = $withJSON.checked;

    doConversion();
});

$compact.addEventListener("change", function () {
    options.compact = $compact.checked;

    doConversion();
});

$sample.addEventListener("click", function () {
    $input.innerText = sample;

    doConversion();
});

// https://www.w3schools.com/xml/xml_examples.asp
// language=XML
const sample = `<note>
    <to>Tove</to>
    <from>Jani</from>
    <heading>Reminder</heading>
    <body>Don't forget me this weekend!</body>
</note>`;

// Go from ./wasm/wasm_exec
const go = new globalThis.Go();
WebAssembly
    .instantiateStreaming(fetch(xmlToGoWasmURL()), go.importObject)
    .then(function (result) {
        go.run(result.instance);
    });
