import { execFileSync } from "node:child_process";
import { rmSync } from "node:fs";

const BUILD_FOLDER = "dist";
const DEPLOY_BRANCH = "gh-pages";

const run = (command, ...args) => execFileSync(command, args, { stdio: "inherit" });

run("git", "checkout", "--orphan", DEPLOY_BRANCH);
run("yarn", "build");
run("git", "config", "user.email", "gorka.guridi@gmail.com");
run("git", "config", "user.name", "Gorka Guridi");
run("git", "--work-tree", BUILD_FOLDER, "add", "--all");
run("git", "--work-tree", BUILD_FOLDER, "commit", "-m", DEPLOY_BRANCH);
run("git", "push", "origin", `HEAD:${DEPLOY_BRANCH}`, "--force");
rmSync(BUILD_FOLDER, { recursive: true });
