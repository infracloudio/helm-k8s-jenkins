node 
{
    docker.withRegistry('https://hub.docker.com/r/harshals/demo-go-app/', 'harshals-Docker-hub') {
    
        git url: "https://github.com/infracloudio/demo-go-app.git", credentialsId: 'harshal-github-creds'
    
        sh "git rev-parse HEAD > .git/commit-id"
        def commit_id = readFile('.git/commit-id').trim()
        println commit_id
    
        stage "build"
        def app = docker.build "demo-go-app"
    
        stage "publish"
        app.push 'master'
        app.push "${commit_id}"
    }
}
