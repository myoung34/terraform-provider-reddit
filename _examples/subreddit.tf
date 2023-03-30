terraform {
  required_providers {
    reddit = {
      source  = "myoung34/reddit"
      version = "0.0.1"
    }
  }
}

provider "reddit" {
  readonly = true
  #id       = "id"
  #secret   = "secret"
  #username = "username"
  #password = "password"
}

data "reddit_subreddit" "golang" {
  name = "golang"
}

output "golang" {
  value = data.reddit_subreddit.golang
}