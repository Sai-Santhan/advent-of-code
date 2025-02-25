import cli from "command-line-args"

export type ProjectorOpts = {
    args?: string[],
    pwd?: string,
    config?: string,
}

export default function getOpts(): ProjectorOpts {
    return cli([
        {
            name: "args",
            defaultOption: true,
            multiple: true,
            type: String
        },
        {
            name: "config",
            alias: "c",
            type: String,
        },
        {
            name: "pwd",
            alias: "p",
            type: String,
        },
    ]) as ProjectorOpts;
}