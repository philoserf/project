workflow "github-pages" {
  on       = "push"

  resolves = [
    "build-deploy"
  ]
}

action "build-deploy" {
  uses    = "BryanSchuetz/jekyll-deploy-gh-pages@master"

  secrets = [
    "GITHUB_TOKEN"
  ]
}
