import "./wasm/wasm_exec"

import highlight from "./highlight";
import {convert, Options} from "./convert";

const $input = document.getElementById("input");
const $output = document.getElementById("output");
const $sample = document.getElementById("sample");
const $inline = document.getElementById("inline") as HTMLInputElement;

const options = new Options($inline.checked);

function doConversion() {
    const result = convert($input.innerText.trim(), options);

    if (result.success) {
        $output.innerHTML = highlight(result.success);
    } else {
        $output.innerHTML = result.error;
    }
}

$input.addEventListener("keyup", doConversion);

$inline.addEventListener("change", function () {
    options.inline = $inline.checked;

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

const xmlToGoWasmURL = "https://xml-to-go.github.io/static/js/wasm/xml-to-go.wasm";
// const xmlToGoWasmURL = "http://localhost:8080/xml-to-go.wasm";

// Go from ./wasm/wasm_exec
const go = new globalThis.Go();
WebAssembly
    .instantiateStreaming(fetch(xmlToGoWasmURL), go.importObject)
    .then(function (result) {
        go.run(result.instance);
    });
