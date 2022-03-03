import fs from "fs";
import { task } from "hardhat/config";
import { HardhatRuntimeEnvironment } from "hardhat/types";

task("go-go-gen", "Builds generate.go file from all json files in path")
  .addParam("in", "relative path of the output files")
  .addParam("out", "relative path of the output files")
  .addParam("pkg", "pkg the go generate command should use")
  .setAction(async (taskArgs, hre: HardhatRuntimeEnvironment) => {
    let outputData = `#! /bin/bash\n`;
    outputData =
      outputData +
      `# DO NOT EDIT THIS FILE! THIS FILE CONTAINS GENERATED CODE\n`;
    outputData = outputData + `# CHANGES MADE TO THIS FILE WILL BE LOST\n`;
    outputData = outputData + `#\n`;
    outputData = outputData + `# npx hardhat go-go-gen \ \n`;
    outputData = outputData + `#    --pkg ${taskArgs.pkg} \ \n`;
    outputData = outputData + `#    --in ${taskArgs.in} \ \n`;
    outputData = outputData + `#    --out ${taskArgs.out}\n`;
    let outpath = taskArgs.out;
    // todo: use a path library for all computed paths below this line
    if (outpath !== "." && outpath[-1] !== "/") {
      outpath = outpath + "/";
    }
    if (outpath === ".") {
      outpath = "";
    }
    fs.readdirSync(taskArgs.in).forEach((file) => {
      let source = file.replace(".json", "");
      outputData =
        outputData +
        `abigen --abi ${taskArgs.in}/${source}.json --pkg ${taskArgs.pkg} --type ${source} --out ${outpath}${source}.go\n`;
    });
    if (!fs.existsSync(`${taskArgs.pkg}`)) {
      fs.mkdirSync(`${taskArgs.pkg}`);
    }
    fs.writeFileSync(`${taskArgs.pkg}/generate.sh`, outputData);
    fs.chmodSync(`${taskArgs.pkg}/generate.sh`, 0o775);
  });
