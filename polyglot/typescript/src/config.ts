import {ProjectorOpts} from "./opts";
import path from "path";

export enum Operation {
    Print,
    Add,
    Remove,
}

export type Config = {
    args: string[],
    operation: Operation,
    config: string,
    pwd: string,
}

function getPwd(opts: ProjectorOpts): string {
    if (opts.pwd) {
        return opts.pwd;
    }
    return process.cwd();
}

function getConfig(opts: ProjectorOpts): string {
    if (opts.config) {
        return opts.config;
    }
    const home = process.env["HOME"]
    const loc = process.env["XDG_CONFIG_HOME"] || home;
    if (!loc) {
        throw new Error("Unable to determine config location");
    }
    if (loc === home) {
        return path.join(loc, ".projector.json");
    }
    return path.join(loc, "projector", "projector.json")
}

function getOperation(opts: ProjectorOpts): Operation {
    if (!opts.args || opts.args.length === 0) {
        return Operation.Print;
    }
    if (opts.args[0] === "add"){
        return Operation.Add;
    }
    if (opts.args[0] === "rm"){
        return Operation.Remove;
    }
    return Operation.Print;
}

function getArgs(opts: ProjectorOpts): string[] {
    if (!opts.args || opts.args.length === 0) {
        return [];
    }

    const operation = getOperation(opts);
    if (operation === Operation.Print) {
          if (opts.args.length > 1){
              throw new Error(`Expected 0 or 1 argument but got ${opts.args.length}.`)
          }
          return opts.args;
    }
    if (operation === Operation.Add) {
          if (opts.args.length != 3){
              throw new Error(`Expected 2 arguments but got ${opts.args.length - 1}.`)
          }
          return opts.args.slice(1);
    }
    if (operation === Operation.Remove) {
          if (opts.args.length != 2){
              throw new Error(`Expected 1 argument but got ${opts.args.length - 1}.`)
          }
          return opts.args.slice(1);
    }
    // return opt
}

export default function config(opts: ProjectorOpts): Config {
    return {
        pwd: getPwd(opts),
        config: getConfig(opts),
        args: getArgs(opts),
        operation: getOperation(opts),
    }
}