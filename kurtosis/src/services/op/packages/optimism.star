GITHUB_URL = "https://github.com/ethereum-optimism/optimism.git"
BRANCH = "tutorials/chain"
ARTIFACT_NAME = "optimism"

def clone(plan):
    output = plan.run_sh(
        image = "alpine/git:latest",
        run = "git clone -b {} {} --depth=1".format(
            BRANCH,
            GITHUB_URL,
        ),
        store = [StoreSpec(src = "./git/optimism", name = ARTIFACT_NAME)],
    )
    return output.files_artifacts[0]
