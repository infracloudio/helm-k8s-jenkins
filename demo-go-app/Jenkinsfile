node 
{
withCredentials([[$class: 'UsernamePasswordMultiBinding', credentialsId: 'docker-hub-creds',
usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD']])
  {

    def appName = "demo-go-app"
    def imageTag = "${USERNAME}/${appName}:${env.BRANCH_NAME}.${env.BUILD_NUMBER}"
    checkout scm 
    stage 'Build image'
      sh("docker build -t ${imageTag} .")

    stage 'Push image to registry'
      sh("docker login -u ${USERNAME} -p ${PASSWORD} ")
      sh("docker push ${imageTag}")
    stage 'Deploy Application'
      sh("kubectl set image deployment/demo-app-${env.BRANCH_NAME} -n prod frontend=${imageTag} ")
  }

}

