# go-with-tests

This repository is dedicated to learning git basics.

### Step 1: Install Git
### Step 2: Configure Git
`git config --global user.name "Your Name" `

`git config --global user.email "youremail@domain.com" `

### Step 3:Clone the Repository
`git clone https:<repo path>`

### Step 4 Create a New Branch

`cd go-with-tests `

`git checkout -b go-basics`

### Step 5: Make Changes and Commit
`git add . `

` git commit -m "Add my changes" `

### Step 6: Push Changes to GitHub
`git push origin go-basics `

### Merge Branch into Main
Once feature branch is up to date and conflicts are resolved, we can merge it into the main branch.

1. Switch to the main branch: `git checkout main`
2. Ensure the main branch is up to date with the remote repository:

`git fetch origin`

`git merge origin/main`

3. Merge feature branch into the main branch: `git merge feature-branch-name` ` git push origin main`

### Create a Pull Request

1. Go to the repository on GitHub.
2. Click on "Pull requests" > "New pull request."
3. Select go-basics as the branch to compare.
4. Review the changes and click "Create pull request."

### Seeing Differences and History
1. To see differences between commits:  `git diff [commit1] [commit2]`
2. To see the history of repository:  `git log `

### Rebasing PRs
Rebasing is used to integrate changes from one branch into another by moving or combining a sequence of commits.

Ensure thebranch is up to date with the main branch by running `git fetch origin` and then `git rebase origin/main`
Force push the changes if necessary with `git push origin go-basics --force`

https://about.gitlab.com/images/press/git-cheat-sheet.pdf







