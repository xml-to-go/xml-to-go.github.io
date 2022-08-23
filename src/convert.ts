class Result {
    constructor(
        public readonly success: string,
        public readonly error: string,
    ) {
    }
}

export class Options {
    constructor(
        public inline: boolean,
        public withJSON: boolean,
        public compact: boolean,
    ) {
    }
}

export function convert(source: string, options: Options): Result {
    return new Result("", "");
}
