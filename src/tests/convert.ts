import * as test from "tape";

import {convert, Options} from "../convert";

test("convert test", (t) => {
    const options = new Options(false);

    function assert(xml: string, goCode: string) {
        t.equal(convert(xml, options).success, goCode);
    }

    // primitives
    {
        // empty
        {
            const xml = "";

            assert(xml, "")
        }
    }

    t.end();
});
