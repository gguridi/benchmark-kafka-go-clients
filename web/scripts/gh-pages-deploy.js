/* eslint-disable no-console */
const execa = require("execa");
const fs = require("fs");
(async () => {
    try {
        await execa("git", ["checkout", "--orphan", "gh-pages"]);
        console.log("Building started...");
        await execa("yarn", ["install"]);
        await execa("yarn", ["build"]);
        const folder = fs.existsSync("dist") ? "dist" : "build";
        await execa("git", ["config", "user.email", "gorka.guridi@gmail.com"]);
        await execa("git", ["config", "user.name", "Gorka Guridi"]);
        await execa("git", ["--work-tree", folder, "add", "--all"]);
        await execa("git", ["--work-tree", folder, "commit", "-m", "gh-pages"]);
        console.log("Pushing to gh-pages...");
        await execa("git", ["push", "origin", "HEAD:gh-pages", "--force"]);
        await execa("rm", ["-r", folder]);
        console.log("Successfully deployed, check your settings");
    } catch (e) {
        console.log(e.message);
        process.exit(1);
    }
})();
