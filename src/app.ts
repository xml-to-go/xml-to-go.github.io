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
